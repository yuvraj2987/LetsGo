package linked_list

func removeDuplicates(head *Node) (head *Node) {
	if head == nil {
		return
	}

	if head.Next == nil {
		// No duplucates in single node list
		return
	}

	cur1 := head

	for (cur1 != nil) || (cur1.Next != nil) {
		cur2 := cur1
		for cur2.Next != nil {
			next2 := cur2.Next
			if cur1.Data == next2.Data {
				// Duplicate found
				cur2.Next = next2.Next
				// Set next2 for GC
				next2.Next = nil
			} else {
				cur2 = cur2.Next
			}
		}
		cur1 = cur1.Next
	}
	return
}
