package linked_list

import (
	"testing"
)

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

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

func TestListToSlice(t *testing.T) {
	var head *Node = nil
	var empty_slice []int
	list_slice := ListToSlice(head)
	// Test Emtpy list
	if !equalSlices(list_slice, empty_slice) {
		t.Errorf("Failed to return emtpy list")
	}
	// Test 1 element list
	Push(&head, 3)
	actual_slice1 := ListToSlice(head)
	expected_slice1 := []int{3}
	if !equalSlices(actual_slice1, expected_slice1) {
		t.Errorf("Actual list = %v expected list = %v", actual_slice1, expected_slice1)
	}

	// Test 2 element list
	Push(&head, 2)
	actual_slice2 := ListToSlice(head)
	expected_slice2 := []int{2, 3}
	if !equalSlices(actual_slice2, expected_slice2) {
		t.Errorf("Actual list = %v expected list = %v", actual_slice2, expected_slice2)
	}
	// Test 3 element list
	Push(&head, 1)
	expected_slice3 := []int{1, 2, 3}
	actual_slice3 := ListToSlice(head)
	if !equalSlices(actual_slice3, expected_slice3) {
		t.Errorf("Actual list = %v expected list = %v", actual_slice3, expected_slice3)
	}

}

func TestAppendList(t *testing.T) {
	var head *Node = nil
	Append(&head, 1)
	// Test1: Append to emtpy
	exp1 := []int{1}
	ans1 := ListToSlice(head)
	if !equalSlices(exp1, ans1) {
		t.Errorf("Actual list = %v; expected list = %v", ans1, exp1)
	}
	// Test2: Append to single node
	Append(&head, 2)
	exp2 := []int{1, 2}
	ans2 := ListToSlice(head)
	if !equalSlices(exp2, ans2) {
		t.Errorf("Actual list = %v; expected list = %v", ans2, exp2)
	}

	// Test3: Append to 2 element list
	Append(&head, 3)
	exp3 := []int{1, 2, 3}
	ans3 := ListToSlice(head)
	if !equalSlices(ans3, exp3) {
		t.Errorf("Actual list = %v; expected list = %v", ans3, exp3)
	}

}