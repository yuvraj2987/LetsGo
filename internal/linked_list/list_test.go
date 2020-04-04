package linked_list

import (
	// "fmt"
	"testing"
)

func TestListPush(t *testing.T) {
	var head *Node = nil
	if Length(head) != 0 {
		t.Errorf("Length mismatch for emtpy list")
	}
	Push(&head, 3)
	if Length(head) != 1 {
		t.Errorf("Length mismatch for single element")
	}
	Push(&head, 2)
	Push(&head, 3)
	if Length(head) != 3 {
		t.Errorf("Length mismatch for 3 elements")
	}
}
