package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Type int

const (
	Env Type = iota
	Str Type = iota
	Color Type = iota
)

type Token struct {
	T Type
	V string
}

var m map[string]Token

func initMap() {
	dir, _ := os.Getwd()
	dir = strings.ReplaceAll(dir,os.Getenv("HOME"), "~")

	hostname, _ := os.ReadFile("/etc/hostname")

	m = make(map[string]Token)
	m["dir"] = Token{T: Str, V: dir}
	m["hostname"] = Token{T: Str, V: string(hostname)[:len(hostname)-1]}
}

func Lexer(data []byte) []Token {
	tokens := []Token{}

	var conf []string
	err := json.Unmarshal(data, &conf)

	if err != nil {
		log.Fatal("malformed config file ", err)
	}

	for _, v := range conf {
		token := Token{}
		if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
			var ok bool
			token, ok = m[v[2:len(v)-1]]
			if !ok {
				token.T = Str
				token.V = v
			}
		} else if strings.HasPrefix(v, "$") {
			token.T = Env
			token.V = v[1:]
		} else if strings.HasPrefix(v, "bg:") || strings.HasPrefix(v, "fg:") || strings.HasPrefix(v, "c:") {
			token.T = Color
			token.V = v
		} else {
			token.T = Str
			token.V = v
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

func main() {

	data, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatal("config file missing", err)
	}
	initMap()
	initColorMap()

	fmt.Println(Parser(Lexer(data)))
}
