package linkedlist

import "fmt"

type LinkedList[T comparable] struct {
	f_node *Node[T]
	c_node *Node[T]
	size   int
}

type Node[T any] struct {
	n_node *Node[T]
	value  T
}

func (linkedList *LinkedList[T]) Append(item T) *Node[T] {

	if linkedList.f_node == nil {
		linkedList.f_node = &Node[T]{nil, item}
		linkedList.c_node = linkedList.f_node

	} else {
		var new_node *Node[T] = &Node[T]{nil, item}
		linkedList.c_node.n_node = new_node
		linkedList.c_node = new_node
	}

	linkedList.size++

	return linkedList.c_node
}

func (linkedList *LinkedList[T]) Remove(index int) bool {

	if linkedList.size == 0 || index < 0 || index >= linkedList.size {
		panic("Invalid Remove")
	} else if index == 0 && linkedList.size > 1 {
		linkedList.f_node = linkedList.f_node.n_node
	} else if index == 0 {
		linkedList.f_node = nil
		linkedList.c_node = nil

	} else {
		counter := 0
		var l_node *Node[T] = nil
		var c_node *Node[T] = linkedList.f_node

		for counter < index {
			l_node = c_node
			c_node = c_node.n_node
			counter++
		}

		if index == (linkedList.size - 1) {
			l_node.n_node = nil
			linkedList.c_node = l_node
			c_node = nil
		} else {
			l_node.n_node = c_node.n_node
			c_node = nil
		}

	}

	linkedList.size--

	return true
}

func (linkedList *LinkedList[T]) Get(index int) T {

	if index < 0 || index >= linkedList.size {
		panic(fmt.Sprintf("Index %d out of bounds (size: %d)", index, linkedList.size))
	}

	current := (*linkedList).f_node

	for i := 0; i < index; i++ {
		current = current.n_node
	}

	return current.value
}

func (linkedList *LinkedList[T]) Size() int {
	return linkedList.size
}

func (linkedList *LinkedList[T]) Prepend(item T) bool {
	var new_node *Node[T] = &Node[T]{nil, item}

	if linkedList.f_node == nil {
		linkedList.f_node = new_node
		linkedList.c_node = linkedList.f_node
	} else {
		new_node.n_node = linkedList.f_node
		linkedList.f_node = new_node
	}

	linkedList.size++

	return true
}

func (linkedList *LinkedList[T]) Insert(index int, item T) bool {

	if index < 0 || index > linkedList.size {
		panic(fmt.Sprintf("index %d out of bounds", index))
	}

	var new_node *Node[T] = &Node[T]{nil, item}

	if index == 0 {
		linkedList.Prepend(item)
	} else if index == linkedList.size {
		linkedList.Append(item)
	} else {
		counter := 0
		var l_node *Node[T] = nil
		var c_node *Node[T] = linkedList.f_node

		for counter < index {
			l_node = c_node
			c_node = c_node.n_node
			counter++
		}

		l_node.n_node = new_node
		new_node.n_node = c_node
	}

	linkedList.size++

	return true
}

func (linkedList *LinkedList[T]) Contains(item T) bool {
	it := linkedList.Iterator()

	for it.HasNext() {
		if it.Next() == item {
			return true
		}
	}

	return false
}

func (linkedList *LinkedList[T]) PrintAll() {

	it := linkedList.Iterator()

	for it.HasNext() {
		fmt.Print(it.Next(), " ")
	}

	fmt.Println()
}

func (linkedList *LinkedList[T]) IsEmpty() bool {
	return linkedList.size > 0
}

func (linkedList *LinkedList[T]) Clone() *LinkedList[T] {

	clone := &LinkedList[T]{}

	it := linkedList.Iterator()

	for it.HasNext() {
		clone.Append(it.Next())
	}

	return clone
}
