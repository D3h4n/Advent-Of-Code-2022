package parser_test

import (
	"day5/pkg/parser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("A command parser should be able to", func() {
	DescribeTable("parse commands",
		func(input string, expectedCommands []parser.Command) {
			// Arrange
			commandParser := parser.NewCommandParser()

			// Act
			commands := commandParser.ParseCommands(input)

			// Assert
			Expect(commands).To(Equal(expectedCommands))
		},
		Entry("with 1 command",
			"move 1 from 1 to 3",
			[]parser.Command{{Amount: 1, From: 1, To: 3}},
		),
		Entry("with 3 commands",
			"move 2 from 4 to 5\nmove 5 from 1 to 2\nmove 3 from 5 to 1",
			[]parser.Command{
				{Amount: 2, From: 4, To: 5},
				{Amount: 5, From: 1, To: 2},
				{Amount: 3, From: 5, To: 1},
			},
		),
	)
})
