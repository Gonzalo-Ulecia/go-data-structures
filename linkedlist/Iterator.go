package linkedlist

type LinkedListIterator[T any] struct {
	c_node *Node[T]
}

func (it *LinkedListIterator[T]) HasNext() bool {
	return it.c_node != nil
}

func (it *LinkedListIterator[T]) Next() T {
	if it.c_node == nil {
		panic("index _ out of bounds")
	}
	val := it.c_node.value
	it.c_node = it.c_node.n_node
	return val
}

func (l *LinkedList[T]) Iterator() *LinkedListIterator[T] {
	return &LinkedListIterator[T]{l.f_node}
}
