package command_parser

import (
	"regexp"
)

func CommandParser(input string) []string {
	args := []string{}
	re := regexp.MustCompile(`"([^"]*)"|'([^']*)'|(\S+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		if match[1] != "" {
			args = append(args, match[1])
		} else if match[2] != "" {
			args = append(args, match[2])
		} else if match[3] != "" {
			args = append(args, match[3])
		}
	}

	return args
}
