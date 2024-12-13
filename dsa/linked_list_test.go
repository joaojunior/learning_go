package dsa

import "testing"

func TestLinkedListAddNodeNextIsNil(t *testing.T){
	linked_list := LinkedList[int]{val: 1}

	new_node := linked_list.add_val(42)

	if linked_list.next != new_node && new_node.val != 42 {
		t.Fatalf("It cannot create a new node on the LinkedList")
	}
}

func TestInverseLinkedListWithOneNode(t *testing.T){
	linked_list := LinkedList[int]{val: 1}

	head := linked_list.invert()

	if linked_list != *head {
		t.Fatalf("It cannot invert the LinkedList with only one node")
	}
}

func TestInverseLinkedListWithTwoNodes(t *testing.T){
	head := LinkedList[int]{val: 1}
	node2 := LinkedList[int]{val: 42}
	head.next = &node2

	new_head := head.invert()

	if *new_head != node2 {
		t.Fatalf("It cannot invert the LinkedList with two nodes")
	}
}

func TestInverseLinkedListWithThreeNodes(t *testing.T){
	head := LinkedList[int]{val: 1}
	node2 := LinkedList[int]{val: 42}
	node3 := LinkedList[int]{val: 84}
	node2.next = &node3
	head.next = &node2

	new_head := head.invert()

	if *new_head != node3 {
		t.Fatalf("It cannot invert the LinkedList with three nodes")
	}
}
