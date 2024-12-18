package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/valid_commands"
	"os"
	"os/exec"
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
				if _, ok = valid_commands.ValidCommandSet[commands[1]]; ok {
					fmt.Printf("%s is a shell builtin\n", commands[1])
				} else {
					if val, err := findInPath(commands[1]); err == nil {
						fmt.Printf("%s is %s\n", commands[1], val)
					} else {
						fmt.Printf("%s: not found\n", commands[1])
					}
				}

			default:
				if val, err := findInPath(commands[0]); err == nil {
					cmd := exec.Command(val, commands[1:]...)
					cmd.Run()
				} else {
					fmt.Printf("%s: not found\n", commands[0])
				}
			}
		}
	}

}

func findInPath(input string) (string, error) {
	var err error
	pathString := os.Getenv("PATH")
	var path []string = strings.Split(pathString, ":")

	for _, dir := range path {
		if _, err := os.Stat(dir + "/" + input); err == nil {
			return dir + "/" + input, err
		}
	}

	err = errors.New("executable not found")
	return "", err
}

func getCommands(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, " ")
}
