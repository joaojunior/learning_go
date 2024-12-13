package dsa

type LinkedList[T any] struct {
	next *LinkedList[T]
	val T
}

func (linked_list *LinkedList[T]) add_val(val T) *LinkedList[T]{
	linked_list.next = &LinkedList[T]{val: val}

	return linked_list.next
}
