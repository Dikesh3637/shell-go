package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/command_parser"
	"github.com/codecrafters-io/shell-starter-go/command_resolver"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		commandSequence := command_parser.CommandParser(input)

		if len(commandSequence) == 0 {
			// Skip empty input
			continue
		}

		command_resolver.ResolveCommand(commandSequence)
	}
}
