package linked_list

import (
	"fmt"
	"testing"
)

func TestListPush(t *testing.T) {
	var head *Node = nil
	fmt.Println(head)
	Print(head)
	Push(&head, 1)
	Print(head)
	Push(&head, 2)
	Push(&head, 3)
	Print(head)
}
