package parser_test

import (
	"day5/pkg/parser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("A Crate Parser should be able to", func() {
	DescribeTable("parse a row of crates",
		func(input string, expectedResult []byte) {
			// Arrange
			crateParser := parser.NewCrateParser()

			// Act
			result := crateParser.ParseRow(input)

			// Assert
			Expect(result).To(Equal(expectedResult))
		},
		Entry("with one crate", "[A]", []byte{'A'}),
		Entry("with three crates", "[A] [B] [C]", []byte{'A', 'B', 'C'}),
		Entry("with missing crates (example 1)", "    [B]", []byte{' ', 'B'}),
		Entry("with missing crates (example 2)", "[A]     [C]", []byte{'A', ' ', 'C'}),
		Entry("with missing crates (example 3)", "[A] [B]    ", []byte{'A', 'B', ' '}),
	)

	DescribeTable("parse several rows of crates and end at row of numbers",
		func(input string, expectedResult [][]byte) {
			// Arrange
			crateParser := parser.NewCrateParser()

			// Act
			result := crateParser.ParseCrates(input)

			// Assert
			Expect(result).To(Equal(expectedResult))
		},
		Entry("with 1 row and stack",
			"[A]\n 1 ",
			[][]byte{{'A'}},
		),
		Entry("with 1 rows and stacks",
			"[A] [B]\n[C] [D]\n 1   2 ",
			[][]byte{{'A', 'B'}, {'C', 'D'}},
		),
		Entry("with 4 rows and 3 stacks",
			"[A]     [B]\n    [C] [D]\n[E] [F]    \n[G] [H] [I]\n 1   2   3 ",
			[][]byte{
				{'A', ' ', 'B'}, {' ', 'C', 'D'}, {'E', 'F', ' '}, {'G', 'H', 'I'},
			},
		),
	)
})
