package main

import (
	"bufio"
	"day5/pkg/parser"
	"day5/pkg/stack"
	"fmt"
	"os"
)

func main() {
	crateParser := parser.NewCrateParser()
	commandParser := parser.NewCommandParser()
	inputFile := "puzzle_input.txt"

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	input, _ := readInput(inputFile)
	crates := crateParser.ParseCrates(input)
	commands := commandParser.ParseCommands(input)

	// ___________
	//   PART 1
	// ___________
	stacks := initialiseStacks(crates)
	performCommands(commands, stacks, false)

	result := ""
	for _, stack := range stacks {
		result += fmt.Sprintf("%c", stack.Pop())
	}
	fmt.Printf("Part 1: %s\n", result)

	// ___________
	//   PART 2
	// ___________
	stacks = initialiseStacks(crates)
	performCommands(commands, stacks, true)

	result = ""
	for _, stack := range stacks {
		result += fmt.Sprintf("%c", stack.Pop())
	}
	fmt.Printf("Part 2: %s\n", result)
}

// readInput takes the name/path of a file and reads its entire content
// as a string.
func readInput(file string) (string, error) {
	f, _ := os.Open(file)
	reader := bufio.NewReader(f)
	return reader.ReadString(0)
}

// initialiseStacks creates and initialises the set of stacks as defined
// in the input file.
func initialiseStacks(crates [][]byte) []stack.Stack {
	stacks := make([]stack.Stack, len(crates[0]))

	for i := range crates {
		for j, crate := range crates[len(crates)-i-1] {
			if crate != ' ' {
				stacks[j].Push(crate)
			}
		}
	}

	return stacks
}

// performCommands iterates through a set of commands in order and modifies
// the set of stacks accordingly. It also accepts a boolean that toggles
// between output for part 1 and part 2 of day 5.
func performCommands(commands []parser.Command, stacks []stack.Stack, isPart2 bool) {
	for _, command := range commands {
		crates := stacks[command.From-1].PopN(command.Amount)

		// reverse crates before pushing for part 2
		if isPart2 {
			for i, j := 0, len(crates)-1; i < j; i, j = i+1, j-1 {
				crates[i], crates[j] = crates[j], crates[i]
			}
		}

		stacks[command.To-1].PushN(crates)
	}
}
