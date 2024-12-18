package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//empty set using a map
	set := make(map[string]struct{})

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := reader.ReadString('\n')

		if _, ok := set[input]; !ok {
			fmt.Printf("%s: command not found", input)
		}
	}
}
