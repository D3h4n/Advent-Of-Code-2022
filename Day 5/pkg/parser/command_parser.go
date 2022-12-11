package parser

import (
	"fmt"
	"strings"
)

type Command struct {
	Amount int
	From   int
	To     int
}

type CommandParser struct{}

func NewCommandParser() *CommandParser {
	return &CommandParser{}
}

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
