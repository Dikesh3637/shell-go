package command_resolver

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/valid_commands"
)

func ResolveCommand(commandSequence []string) {
	command := commandSequence[0]
	command_args := []string{}
	redirect_args := []string{}
	originalStdout := os.Stdout

	redirect_idx, _ := findRedirectArgs(commandSequence)

	if redirect_idx != -1 {
		command_args = commandSequence[1:redirect_idx]
		redirect_args = commandSequence[redirect_idx:]
		redirectStdout(redirect_args[1], redirect_args[0])

	} else {
		command_args = commandSequence[1:]
	}

	switch command {

	case "exit":
		os.Exit(0)

	case "echo":
		for _, echoString := range command_args {
			fmt.Print(echoString)
			fmt.Print(" ")
		}
		fmt.Printf("\n")

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
		fmt.Print(pwd)
		fmt.Printf("\n")

	case "cd":
		if len(commandSequence) == 1 {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error fetching working directory: %s", err)
			} else {
				fmt.Print(pwd)
			}
		} else {
			if commandSequence[1] == "~" {
				commandSequence[1] = os.Getenv("HOME")
			}
			if err := os.Chdir(commandSequence[1]); err != nil {
				fmt.Printf("cd: %s: %s\n", commandSequence[1], "No such file or directory")
			}
		}

	default:
		_, err := findExecutablePath(commandSequence[0])
		if err != nil {
			fmt.Printf("%s: command not found\n", commandSequence[0])
			return
		}
		cmd := exec.Command(commandSequence[0], command_args...)
		cmd.Env = os.Environ()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	os.Stdout = originalStdout // reset os.Stdout

}

func findExecutablePath(target string) (string, error) {
	pathEnv, isSet := os.LookupEnv("PATH")
	if !isSet {
		return "", errors.New("PATH variable is not set")
	}

	paths := strings.Split(pathEnv, string(os.PathListSeparator))
	for _, dir := range paths {
		fullPath := filepath.Join(dir, target)
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, nil
		}
	}

	return "", errors.New("executable file not found in PATH")
}

func findRedirectArgs(arr []string) (int, string) {
	for i, element := range arr {
		if element == ">" || element == "1>" || element == "2>" {
			return i, element
		}
	}
	return -1, ""
}

func redirectStdout(targetPath string, redirectOp string) error {
	// Create parent directories if they don't exist
	ensureDir(targetPath)

	// Open file for writing, create if doesn't exist, truncate if exists
	file, err := os.OpenFile(targetPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}

	if redirectOp == "2>" {
		// Set stderr to our new file
		os.Stderr = file
	} else {
		//Set stdout to our new file
		os.Stdout = file
	}
	return nil
}

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.ModePerm)
		if merr != nil {
			panic(merr)
		}
	}
}
