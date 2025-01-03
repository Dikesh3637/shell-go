package command_parser

import (
	"strings"
)

func CommandParser(input string) []string {
	input = strings.TrimSpace(input)
	args := []string{}

	doubleQuotedSpecialCharacters := map[rune]string{
		'$':  "",
		'\\': "",
		'"':  "",
	}

	const (
		isSingleQouted = iota
		isDoubleQouted
		isEscaped
	)

	currentState := isEscaped
	skipNext := false
	buffer := ""

	for i, arg := range input {
		if skipNext {
			skipNext = false
			continue
		}

		if i == 0 {
			if arg == '"' {
				currentState = isDoubleQouted
				continue
			} else if arg == '\'' {
				currentState = isSingleQouted
				continue
			}
		}

		switch arg {
		case '"':
			if currentState == isEscaped {
				currentState = isDoubleQouted
				buffer += string(input[i+1])
				skipNext = true
			} else if currentState == isDoubleQouted {
				currentState = isEscaped
			} else {
				buffer += string(arg)
			}
		case '\'':
			if currentState == isEscaped {
				currentState = isSingleQouted
				buffer += string(input[i+1])
				skipNext = true
			} else if currentState == isSingleQouted {
				currentState = isEscaped
			} else {
				buffer += string(arg)
			}
		case '\\':
			if currentState == isEscaped {
				buffer += string(input[i+1])
				skipNext = true
			} else if currentState == isDoubleQouted {
				_, ok := doubleQuotedSpecialCharacters[rune(input[i+1])]
				if ok {
					buffer += string(input[i+1])
					skipNext = true
				} else {
					buffer += string(arg)
				}
			} else if currentState == isSingleQouted {
				buffer += string(arg)
			}
		case ' ':
			if currentState == isEscaped {
				args = append(args, buffer)
				buffer = ""
			} else {
				buffer += string(arg)
			}

		default:
			buffer += string(arg)
		}
	}

	if len(buffer) > 0 {
		args = append(args, buffer)
	}

	res := []string{}

	for _, arg := range args {
		if len(arg) > 0 {
			res = append(res, arg)
		}
	}

	return res
}
