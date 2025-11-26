package main

import (
	"fmt"
	"strconv"
)

var color_map map[string]string

const FOREGROUND_ESCAPE = "\033[38;2;%d;%d;%dm"
const BACKGROUND_ESCAPE = "\033[48;2;%d;%d;%dm"

func initColorMap() {
	color_map = make(map[string]string)
	color_map["c:reset"] = "\x1b[0m"

	color_map["c:bold"] = "\x1b[1m"
	color_map["c:dim"] = "\x1b[2m"
	color_map["c:italic"] = "\x1b[3m"
	color_map["c:underline"] = "\x1b[4m"
	color_map["c:blink"] = "\x1b[5m"
	color_map["c:reverse"] = "\x1b[7m"
	color_map["c:hidden"] = "\x1b[8m"
	color_map["c:strike"] = "\x1b[9m"

	color_map["fg:black"] = "\x1b[30m"
	color_map["fg:red"] = "\x1b[31m"
	color_map["fg:green"] = "\x1b[32m"
	color_map["fg:yellow"] = "\x1b[33m"
	color_map["fg:blue"] = "\x1b[34m"
	color_map["fg:magenta"] = "\x1b[35m"
	color_map["fg:cyan"] = "\x1b[36m"
	color_map["fg:white"] = "\x1b[37m"

	color_map["fg:bright_black"] = "\x1b[90m"
	color_map["fg:bright_red"] = "\x1b[91m"
	color_map["fg:bright_green"] = "\x1b[92m"
	color_map["fg:bright_yellow"] = "\x1b[93m"
	color_map["fg:bright_blue"] = "\x1b[94m"
	color_map["fg:bright_magenta"] = "\x1b[95m"
	color_map["fg:bright_cyan"] = "\x1b[96m"
	color_map["fg:bright_white"] = "\x1b[97m"

	color_map["bg:black"] = "\x1b[40m"
	color_map["bg:red"] = "\x1b[41m"
	color_map["bg:green"] = "\x1b[42m"
	color_map["bg:yellow"] = "\x1b[43m"
	color_map["bg:blue"] = "\x1b[44m"
	color_map["bg:magenta"] = "\x1b[45m"
	color_map["bg:cyan"] = "\x1b[46m"
	color_map["bg:white"] = "\x1b[47m"

	color_map["bg:bright_black"] = "\x1b[100m"
	color_map["bg:bright_red"] = "\x1b[101m"
	color_map["bg:bright_green"] = "\x1b[102m"
	color_map["bg:bright_yellow"] = "\x1b[103m"
	color_map["bg:bright_blue"] = "\x1b[104m"
	color_map["bg:bright_magenta"] = "\x1b[105m"
	color_map["bg:bright_cyan"] = "\x1b[106m"
	color_map["bg:bright_white"] = "\x1b[107m"
}

func ParseColor(content string) string {
	color, ok := color_map[content]

	prefix := content[:2]

	if !ok && len(content) == 10 && content[3] == '#' {
		r, g, b, err := hexToRGB(content[4:])
		if err != nil {
			fmt.Println(err)
			return color
		}

		switch prefix {
		case "fg":
			color = fmt.Sprintf(FOREGROUND_ESCAPE, r, g, b)
		case "bg":
			color = fmt.Sprintf(BACKGROUND_ESCAPE, r, g, b)
		}
	}

	return color
}

func hexToRGB(hex string) (int, int, int, error) {
	r, err := strconv.ParseInt(hex[0:2], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 0)
	if err != nil {
		return 0, 0, 0, err
	}

	return int(r), int(g), int(b), nil
}
