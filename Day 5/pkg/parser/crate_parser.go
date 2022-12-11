package parser

import (
	"fmt"
	"strings"
)

type CrateParser struct{}

func NewCrateParser() *CrateParser {
	return &CrateParser{}
}

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
