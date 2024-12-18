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
		var command []string = getCommands(input)

		if _, ok := valid_commands.ValidCommandSet[command[0]]; !ok {
			fmt.Printf("%s: command not found\n", input)
		} else {
			switch command[0] {
			case "exit":
				return
			}
		}
	}

}

func getCommands(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, " ")
}
