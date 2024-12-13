package dsa

import "testing"

func TestLinkedListAddNodeNextIsNil(t *testing.T){
	linked_list := LinkedList[int]{val: 1}

	new_node := linked_list.add_val(42)

	if linked_list.next != new_node && new_node.val != 42 {
		t.Fatalf("It cannot create a new node on the LinkedList")
	}
}
