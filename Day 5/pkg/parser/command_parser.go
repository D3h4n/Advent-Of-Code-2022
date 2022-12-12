package parser

import (
	"fmt"
	"strings"
)

// Command represents a move command.
type Command struct {
	// Amount of crates to move.
	Amount int
	// From represents the source stack.
	From int
	// To represents the target stack.
	To int
}

// CommandParser is reponsible for parsing the list of move commands.
type CommandParser struct{}

// NewCommandParser creates a CommandParser, which is reponsible for parsing
// the list of move commands.
func NewCommandParser() *CommandParser {
	return &CommandParser{}
}

// ParseCommands from an input string, skipping the crate definitions.
func (this *CommandParser) ParseCommands(input string) []Command {
	commands := []Command{}

	for _, str := range strings.Split(input, "\n") {
		command := Command{}

		if count, err := readCommand(str, &command); count != 3 || err != nil {
			continue
		}

		commands = append(commands, command)
	}

	return commands
}

func readCommand(str string, command *Command) (int, error) {
	return fmt.Sscanf(str, "move %d from %d to %d",
		&command.Amount, &command.From, &command.To)
}
