// Package stack provides a basic Stack implementation.
package stack

// Stack stores items in a First-In-Last-Out order.
type Stack struct {
	data    []interface{}
	size    int
	maxSize int
}

// NewStack creates a new Stack.
func NewStack() *Stack {
	return &Stack{}
}

// Push a new item to the stack.
func (this *Stack) Push(value interface{}) {
	if this.size == this.maxSize {
		this.data = append(this.data, value)
		this.maxSize++
	} else {
		this.data[this.size] = value
	}

	this.size++
}

// PushN items to the stack.
func (this *Stack) PushN(items []interface{}) {
	for _, item := range items {
		this.Push(item)
	}
}

// Pop an item from the stack.
func (this *Stack) Pop() interface{} {
	if this.size == 0 {
		return nil
	}

	this.size--
	return this.data[this.size]
}

// PopN items from the stack.
func (this *Stack) PopN(n int) []interface{} {
	items := []interface{}{}

	for i := 0; i < n; i++ {
		items = append(items, this.Pop())
	}

	return items
}

// Size of the stack
func (this *Stack) Size() int {
	return this.size
}
