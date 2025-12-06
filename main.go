package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Type int

const (
	Env Type = iota
	Str
	Color
)

type Token struct {
	T Type
	V string
}

var m map[string]Token

func initMap() {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir, os.Getenv("HOME"), "~")

	hostname, _ := os.ReadFile("/etc/hostname")

	m = make(map[string]Token)
	m["dir"] = Token{T: Str, V: dir}
	m["hostname"] = Token{T: Str, V: string(hostname)[:len(hostname)-1]}
	m["git_branch"] = Token{T: Str, V: GitBranch()}
}

func Converter(data []byte) []any {
	var conf []any
	err := json.Unmarshal(data, &conf)

	if err != nil {
		fmt.Println("malformed config file ", err)
		panic(err)
	}

	return conf
}

func Lexer(conf []any) []Token {

	tokens := []Token{}

	for _, confVal := range conf {
		token := Token{}

		switch val := confVal.(type) {
		case string:
			if strings.HasPrefix(val, "${") && strings.HasSuffix(val, "}") {
				var ok bool
				token, ok = m[val[2:len(val)-1]]
				if !ok {
					token.T = Str
					token.V = val
				}
			} else if strings.HasPrefix(val, "$") {
				token.T = Env
				token.V = val[1:]
			} else if strings.HasPrefix(val, "bg:") || strings.HasPrefix(val, "fg:") || strings.HasPrefix(val, "c:") {
				token.T = Color
				token.V = val
			} else if strings.HasPrefix(val, "exec:") {
				token.T = Str
				token.V = Execute(val[5:])
			} else {
				token.T = Str
				token.V = val
			}
		case map[string]any:
			for k, v := range val {
				switch k {
				case "git_status_noclean":
					if arr, ok := v.([]any); ok {
						content := GitStatus_NoClean(arr)
						if content != "" {
							token.T = Str
							token.V = content
						}
					}
				case "git_status_clean":
					if arr, ok := v.([]any); ok {
						content := GitStatus_Clean(arr)
						if content != "" {
							token.T = Str
							token.V = content
						}
					}
				}
			}
		}

		tokens = append(tokens, token)
	}

	return tokens
}

func Parser(tokens []Token) string {
	prompt := ""
	for _, t := range tokens {
		switch t.T {
		case Env:
			prompt += os.Getenv(t.V)
		case Str:
			prompt += t.V
		case Color:
			prompt += ParseColor(t.V)
		}
	}
	prompt += ParseColor("c:reset")

	return prompt
}

const CONFIGFILE_NAME = "grompt.json"

func main() {
	homePath, err := os.UserHomeDir()

	if err != nil {
		fmt.Println("Home dir error")
		panic(err)
	}

	configPath := fmt.Sprintf("%s/.config/%s", homePath, CONFIGFILE_NAME)
	customPath := os.Getenv("CONFIG_PATH")
	if customPath != "" {
		configPath = customPath
	}

	_, err = os.Stat(configPath)

	if err != nil && !os.IsExist(err) {

		content := `
			[
				"$USER",
				"@",
				"fg:green",
				"${dir}",
				":",
				" "
			]
		`
		os.WriteFile(configPath, []byte(content), 0644)
		fmt.Println("config file generated in " + configPath)
	}

	content, err := os.ReadFile(configPath)

	if err != nil {
		fmt.Println("config file missing", err)
		panic(err)
	}

	initMap()
	initColorMap()

	data := Converter(content)
	fmt.Println(Parser(Lexer(data)))
}
