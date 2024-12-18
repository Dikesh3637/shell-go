package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//empty set using a map
	set := make(map[string]struct{})
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if _, ok := set[input]; !ok {
			fmt.Printf("%s: command not found\n", input)
		} else {
			switch input {
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
