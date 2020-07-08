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
		t.Errorf("test1 failed")
	}

}

func TestRemoveDuplicates(t *testing.T) {
	var head *Node = nil
	// Test1: Emtpy list
	removeDuplicates(head)

	if head != nil {
		t.Errorf("Test1: Failed")
	}

	// Test2: Single element list
	Append(&head, 1)
	removeDuplicates(head)
	exp1 := []int{1}
	ans1 := ListToSlice(head)
	if !equalSlices(exp1, ans1) {
		t.Errorf("Actual list = %v; expected list = %v", ans1, exp1)
	}

	// Test3: 3 elements
	Append(&head, 1)
	Append(&head, 2)
	before_exp2 := []int{1, 1, 2}
	before_list2 := ListToSlice(head)
	if !equalSlices(before_exp2, before_list2) {
		t.Errorf("Actual list = %v; expected list = %v", before_list2, before_exp2)
	}
	removeDuplicates(head)
	after_exp2 := []int{1, 2}
	after_list2 := ListToSlice(head)
	if !equalSlices(after_exp2, after_list2) {
		t.Errorf("Actual list = %v; expected list = %v", after_list2, after_exp2)
	}

	// Test4: 3 elements last element is duplicate
	head2 := Node{3, nil}
	headRef2 := &head2
	Append(&headRef2, 2)
	Append(&headRef2, 2)
	before_exp4 := []int{3, 2, 2}
	before_list4 := ListToSlice(headRef2)
	if !equalSlices(before_exp4, before_list4) {
		t.Errorf("Error while creating test list 4")
		t.Errorf("Actual list = %v; expected list = %v", before_list4, before_exp4)
	}
	removeDuplicates(headRef2)
	after_exp4 := []int{3, 2}
	after_list4 := ListToSlice(headRef2)

	if !equalSlices(after_exp4, after_list4) {
		t.Errorf("Actual list = %v; expected list = %v", after_list4, after_exp4)
	}
	// Test5: 3 elements 1st and last is duplicate
	Append(&headRef2, 3)
	before_exp5 := []int{3, 2, 3}
	before_list5 := ListToSlice(headRef2)
	if !equalSlices(before_exp5, before_list5) {
		t.Errorf("Error while creating test list 5")
		t.Errorf("Actual list = %v; expected list = %v", before_list5, before_exp5)
	}
	removeDuplicates(headRef2)
	after_exp5 := []int{3, 2}
	after_list5 := ListToSlice(headRef2)

	if !equalSlices(after_exp5, after_list5) {
		t.Errorf("Actual list = %v; expected list = %v", after_list5, after_exp5)
	}

	// Test6: No removal
	Append(&headRef2, 1)
	Append(&headRef2, 5)
	before_exp6 := []int{3, 2, 1, 5}
	before_list6 := ListToSlice(headRef2)
	if !equalSlices(before_exp6, before_list6) {
		t.Errorf("Error while creating test list 6")
		t.Errorf("Actual list = %v; expected list = %v", before_list6, before_exp6)
	}
	removeDuplicates(headRef2)
	after_list6 := ListToSlice(headRef2)
	if !equalSlices(before_exp6, after_list6) {
		t.Errorf("Actual list = %v; expected list = %v", before_list6, before_exp6)
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
