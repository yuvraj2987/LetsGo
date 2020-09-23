package stack

type Stack interface {
	Push(interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmtpy() bool
}

type intNode struct {
	next  *intNode
	value int
}

type IntStack struct {
	top *intNode
}

func New() *IntStack {
	return &IntStack{nil}
}

func (this *IntStack) Push(val int) {
	nd := intNode{this.top, val}
	this.top = &nd
}

func (this *IntStack) Pop() int {
	ret := this.top
	this.top = ret.next
	return ret.value
}

func (this *IntStack) Peek() int {
	return this.top.value
}
