package parser

import (
	"fmt"
	"strings"
)

// CrateParser is responsible for parsing stacks of crates.
type CrateParser struct{}

// NewCrateParser creates a crate parser, which is reponsible
// for parsing the various stacks of crates from an input string.
func NewCrateParser() *CrateParser {
	return &CrateParser{}
}

// ParseRow of crates across several stacks.
func (this *CrateParser) ParseRow(row string) []byte {
	crates := []byte{}

	for i := 0; i < len(row); i += 4 {
		if row[i] == ' ' {
			crates = append(crates, 0)
		} else if row[i] == '[' {
			crates = append(crates, row[i+1])
		} else {
			panic(fmt.Sprintf("Unreachable case at i = %d.\nThe row may be incorrectly formatted.", i))
		}
	}

	return crates
}

// ParseCrates from input string, stopping at row of stack numbers.
func (this *CrateParser) ParseCrates(input string) ([][]byte, int) {
	crates := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		if row[1] > '0' && row[1] < '9' {
			return crates, len(crates[0])
		}

		crates = append(crates, this.ParseRow(row))
	}

	panic("Unreachable Code. The input string may be incorrectly formatted.")
}
