package command_parser_test

import (
	"fmt"
	"github.com/codecrafters-io/shell-starter-go/command_parser"
	"testing"
)

func TestCommandParser(t *testing.T) {
	result := command_parser.CommandParser(`echo 'hello world'`)
	expected := []string{"echo", "hello world"}
	fmt.Print("result", result)
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
