package dsa

type LinkedList[T any] struct {
	next *LinkedList[T]
	val T
}

func (linked_list *LinkedList[T]) add_val(val T) *LinkedList[T]{
	linked_list.next = &LinkedList[T]{val: val}

	return linked_list.next
}

func (linked_list *LinkedList[T]) invert() * LinkedList[T]{
	temp := linked_list
	previous := linked_list.next
	current := linked_list
	for ; current != nil ; {
		temp = current.next
		current.next = previous
		previous = current
		current = temp
	}
	return previous
}
