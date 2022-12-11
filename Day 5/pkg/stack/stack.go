package stack

type Stack struct {
	data []interface{}
	tail int
	size int
}

func NewStack() *Stack {
	return &Stack{}
}

func (this *Stack) Push(value interface{}) {
	if this.tail == this.size {
		this.data = append(this.data, value)
		this.size++
	} else {
		this.data[this.tail] = value
	}

	this.tail++
}

func (this *Stack) PushN(items []interface{}) {
	for _, item := range items {
		this.Push(item)
	}
}

func (this *Stack) Size() int {
	return this.tail
}

func (this *Stack) Pop() interface{} {
	if this.tail == 0 {
		return nil
	}

	this.tail--
	return this.data[this.tail]
}

func (this *Stack) PopN(n int) []interface{} {
	items := []interface{}{}

	for i := 0; i < n; i++ {
		items = append(items, this.Pop())
	}

	return items
}
