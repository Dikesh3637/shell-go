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

// func getCommands(input string) []string {
// 	input = strings.TrimSpace(input)

// 	doubleQuoted := []string{}
// 	singleQuoted := []string{}

// 	for i,char := range(input){
// 		if char == '"' {
// 			for j := i+1; j < len(input); j++ {
//                 if input[j] == '"' {
//                     doubleQuoted = append(doubleQuoted, input[i+1:j])
//                     i = j+1
//                     break
//                 }
//             }
// 		}
// 		if char == '\'' {
//             for j := i+1; j < len(input); j++ {
//                 if input[j] == '\'' {
//                     singleQuoted = append(singleQuoted, input[i+1:j])
//                     i = j+1
//                     break
//                 }
//             }
//         }

// 	}

// 	}

// }
