package command_parser

import "strings"

func isSpace(char rune) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func CommandParser(input string) []string {
	input = strings.TrimSpace(input)
	args := []string{}

	const (
		isSingleQouted = iota
		isDoubleQouted
		isEscaped
	)

	currentState := isEscaped
	skipNext := false
	buff := ""

	for i, arg := range input {
		if skipNext {
			skipNext = false
			continue
		}

		switch arg {
		case '\\':
			if currentState == isDoubleQouted {
				buff += string(arg)

			} else if currentState == isSingleQouted {
				buff += string(arg)

			} else {
				buff += string(input[i+1])
				skipNext = true
			}
		case '"':
			if currentState == isEscaped {
				args = append(args, buff)
				buff = ""
				currentState = isDoubleQouted
			} else if currentState == isDoubleQouted {
				args = append(args, buff)
				buff = ""
				currentState = isEscaped
			} else {
				buff += string(arg)
			}
		case '\'':
			if currentState == isDoubleQouted {
				buff += string(arg)
			} else if currentState == isSingleQouted {
				args = append(args, buff)
				buff = ""
				currentState = isEscaped
			} else {
				args = append(args, buff)
				buff = ""
				currentState = isSingleQouted
			}
		default:
			if arg == ' ' {
				if currentState == isEscaped {
					args = append(args, buff)
					buff = ""
					break
				}
			}

			buff += string(arg)
		}
	}

	args = append(args, buff)

	res := []string{}

	for _, arg := range args {
		if len(arg) > 0 {
			res = append(res, arg)
		}
	}

	return res
}
