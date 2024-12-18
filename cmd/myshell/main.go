package main

import (
	"bufio"
	"fmt"
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
		commands := getCommands(input)

		if len(commands) == 0 {
			// Skip empty input
			continue
		}

		switch commands[0] {
		case "exit":
			return

		case "echo":
			echoString := strings.Join(commands[1:], " ")
			fmt.Println(echoString)

		case "type":
			if _, ok := valid_commands.ValidCommandSet[commands[1]]; ok {
				fmt.Printf("%s is a shell builtin\n", commands[1])
			} else {
				if val, err := exec.LookPath(commands[1]); err == nil {
					fmt.Printf("%s is %s\n", commands[1], val)
				} else {
					fmt.Printf("%s: not found\n", commands[1])
				}
			}

		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)

		case "cd":
			pwd, _ := os.Getwd()
			if len(commands) == 1 {
				fmt.Println(pwd)
			} else {
				if commands[1] == "~" {
					commands[1] = os.Getenv("HOME")
				}
				if err := os.Chdir(commands[1]); err != nil {
					fmt.Printf("cd: %s: %s\n", commands[1], "No such file or directory")
				}
			}

		default:
			cmd := exec.Command(commands[0], commands[1:]...)
			cmdOutput, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("%s: %s\n", commands[0], "command not found")
			} else {
				fmt.Print(string(cmdOutput))
			}
		}
	}
}

func getCommands(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, " ")
}
