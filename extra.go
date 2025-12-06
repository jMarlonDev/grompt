package main

import (
	"os/exec"
	"strings"
)

func GitBranch() string {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	content, _ := cmd.Output()

	return strings.Trim(string(content), "\n")
}

func GitStatus_NoClean(data []any) string {

	cmd := exec.Command("git", "status", "--porcelain")

	content, _ := cmd.Output()

	output := ""

	if len(content) > 1 {
		output = Parser(Lexer(data))
	}

	return output
}

func GitStatus_Clean(data []any) string {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()

	if err != nil {
		return ""
	}

	cmd = exec.Command("git", "status", "--porcelain")

	content, _ := cmd.Output()

	output := ""

	if len(content) < 1 {
		output = Parser(Lexer(data))
	}

	return output
}

func Execute(cmd string) string {
	cmdParts := strings.Split(cmd, " ")
	command := exec.Command(cmdParts[0], cmdParts[1:]...)
	output, _ := command.Output()
	return strings.Trim(string(output), "\n")
}
