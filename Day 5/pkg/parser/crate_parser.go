package parser

import (
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

	for i := 1; i < len(row); i += 4 {
		crates = append(crates, row[i])
	}

	return crates
}

// ParseCrates from input string, stopping at row of stack numbers.
func (this *CrateParser) ParseCrates(input string) [][]byte {
	crates := [][]byte{}

	for _, row := range strings.Split(input, "\n") {
		if is_ascii_numeric(row[1]) {
			return crates
		}

		crates = append(crates, this.ParseRow(row))
	}

	panic("Unreachable Code. The input string may be incorrectly formatted.")
}

func is_ascii_numeric(c byte) bool {
	return c >= '0' && c <= '9'
}
