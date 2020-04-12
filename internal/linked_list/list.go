package linked_list

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func Push(headRef **Node, data int) {
	new_node := Node{data, *headRef}
	*headRef = &new_node
}

func Append(headRef **Node, data int) {
	dummy_head := Node{0, *headRef}
	var tail *Node = &dummy_head
	for tail.Next != nil {
		tail = tail.Next
	}
	Push(&(tail.Next), data)
	*headRef = dummy_head.Next
}

func Print(head *Node) {
	if head == nil {
		fmt.Println("Emtpy List")
	}
	cur := head
	for cur != nil {
		fmt.Printf("%d->", cur.Data)
		cur = cur.Next
	}
	fmt.Println()
}

func Length(head *Node) int {
	cur := head
	count := 0
	for cur != nil {
		cur = cur.Next
		count += 1
	}

	return count
}

func ListToSlice(head *Node) []int {
	var lice []int
	if head == nil {
		return lice
	}
	lice = make([]int, 0, 1)
	lice = append(lice, head.Data)
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		lice = append(lice, cur.Data)
	}
	return lice
}
