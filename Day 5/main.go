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
	isPart2 := false

	if len(os.Args) > 1 {
		inputFile = os.Args[1]
	}

	input, _ := readInput(inputFile)
	crates, numStacks := crateParser.ParseCrates(input)
	commands := commandParser.ParseCommands(input)

	stacks := initialiseStacks(crates, numStacks)
	performCommands(commands, stacks, isPart2)

	fmt.Printf("stacks: %v\n", stacks)

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Pop())
	}
	println()
}

func readInput(filename string) (string, error) {
	f, _ := os.Open(filename)
	reader := bufio.NewReader(f)
	return reader.ReadString(0)
}

func initialiseStacks(crates [][]byte, numStacks int) []stack.Stack {
	stacks := make([]stack.Stack, numStacks)

	for i := range crates {
		for i, crate := range crates[len(crates)-i-1] {
			if crate != 0 {
				stacks[i].Push(crate)
			}
		}
	}

	return stacks
}

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
