package stack_test

import (
	_stack "day5/pkg/stack"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStack(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Stack Test Suite")
}

var _ = DescribeTable("Can store values",
	func(values []string, size int) {
		// Arrange
		stack := _stack.NewStack()

		// Act
		for _, value := range values {
			stack.Push(value)
		}

		// Assert
		Expect(stack.Size()).To(Equal(size))
	},
	Entry("With 1 value", []string{"A"}, 1),
	Entry("With 3 values", []string{"A", "Gertrude", "The Quick Brown Fox Jumps Over the Lazy Dog"}, 3),
)

var _ = DescribeTable("pops values in correct order",
	func(values []string) {
		// Arrange
		stack := _stack.NewStack()

		for _, value := range values {
			stack.Push(value)
		}

		for i := range values {
			// Act
			output := stack.Pop()

			// Assert
			Expect(output).To(Equal(values[len(values)-i-1]))
		}
	},
	Entry("With 1 value", []string{"A"}),
	Entry("With 3 values", []string{"A", "Gertrude", "The Quick Brown Fox Jumps Over the Lazy Dog"}),
)

var _ = It("returns nil when popping an empty stack", func() {
	// Arrange
	stack := _stack.NewStack()

	// Act
	value := stack.Pop()

	// Assert
	Expect(value).To(BeNil())
})

var _ = DescribeTable("pushes n items",
	func(items []interface{}) {
		// Arrange
		stack := _stack.NewStack()

		// Act
		stack.PushN(items)

		// Assert
		Expect(stack.Size()).To(Equal(len(items)))
	},
	Entry("1 item", []interface{}{1, 2, 3}),
	Entry("3 items", []interface{}{3, 5, 8, 6, 7}),
)

var _ = DescribeTable("pops n items",
	func(items []interface{}, n int, expectedResult []interface{}) {
		// Arrange
		stack := _stack.NewStack()
		stack.PushN(items)

		// Act
		result := stack.PopN(n)

		// Assert
		Expect(result).To(Equal(expectedResult))
		Expect(stack.Size()).To(Equal(len(items) - n))
	},
	Entry("1 item", []interface{}{1, 2, 3}, 1, []interface{}{3}),
	Entry("3 items", []interface{}{3, 5, 8, 6, 7}, 3, []interface{}{7, 6, 8}),
)
