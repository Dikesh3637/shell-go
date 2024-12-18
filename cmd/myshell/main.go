package main

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/valid_commands"
	"os"
	"strings"
)

func main() {
	//empty set using a map
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		var commands []string = getCommands(input)

		if _, ok := valid_commands.ValidCommandSet[commands[0]]; !ok {
			fmt.Printf("%s: command not found\n", commands[0])
		} else {
			switch commands[0] {
			case "exit":
				return
			case "echo":
				var echoString string = strings.Join(commands[1:], " ")
				fmt.Println(echoString)
			case "type":
				if _, ok = valid_commands.ValidCommandSet[commands[1]]; !ok {
					fmt.Printf("%s: not found\n", commands[1])
				} else {
					fmt.Printf("%s is a shell builtin\n", commands[1])
				}

			default:
				fmt.Printf("%s: not found\n", commands[0])
			}
		}
	}

}

func getCommands(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, " ")
}
