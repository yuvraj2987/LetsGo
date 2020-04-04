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
