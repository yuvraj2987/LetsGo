package linked_list

import (
	"fmt"
	"testing"
)

func TestListPush(t *testing.T) {
	root := new(Node)
	fmt.Println("Root Ref: ", root)
	Print(root)
	Push(&root, 1)
	Print(root)
	Push(&root, 2)
	Push(&root, 3)
	Print(root)
}
