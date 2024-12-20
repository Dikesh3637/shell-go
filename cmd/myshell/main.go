package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/command_parser"
	"github.com/codecrafters-io/shell-starter-go/valid_commands"
	"os"
	"os/exec"
	"strings"
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

		switch commandSequence[0] {
		case "exit":
			return

		case "echo":
			echoString := strings.Join(commandSequence[1:], " ")
			fmt.Println(echoString)

		case "type":
			if _, ok := valid_commands.ValidCommandSet[commandSequence[1]]; ok {
				fmt.Printf("%s is a shell builtin\n", commandSequence[1])
			} else {
				if val, err := exec.LookPath(commandSequence[1]); err == nil {
					fmt.Printf("%s is %s\n", commandSequence[1], val)
				} else {
					fmt.Printf("%s: not found\n", commandSequence[1])
				}
			}

		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)

		case "cd":
			pwd, _ := os.Getwd()
			if len(commandSequence) == 1 {
				fmt.Println(pwd)
			} else {
				if commandSequence[1] == "~" {
					commandSequence[1] = os.Getenv("HOME")
				}
				if err := os.Chdir(commandSequence[1]); err != nil {
					fmt.Printf("cd: %s: %s\n", commandSequence[1], "No such file or directory")
				}
			}

		default:
			cmd := exec.Command(commandSequence[0], commandSequence[1:]...)
			cmdOutput, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s: %s\n", commandSequence[0], "command not found")
			} else {
				fmt.Print(string(cmdOutput))
			}
		}
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
