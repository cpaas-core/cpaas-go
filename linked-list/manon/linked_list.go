package linkedlist

import "errors"

// Node type
type Node struct {
	Value      interface{}
	prev, next *Node
}

type List struct {
	head, tail *Node
}

func NewList(elements ...interface{}) *List {
	newList := &List{}
	for _, element := range elements {
		newList.Push(element)
	}
	return newList
}

// Next return the following node of the linked-list towards the tail
func (n *Node) Next() *Node {
	return n.next
}

// Prev return the following node of the linked-list towards the head
func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift insert value at the front of the list. Same as push in the reversed linked-list
func (l *List) Unshift(v interface{}) {
	l.Reverse()
	l.Push(v)
	l.Reverse()
}

// Push insert a value at the back of the list
func (l *List) Push(v interface{}) {
	// Create new node
	newNode := &Node{v, l.tail, nil}
	// Include the new node in the tail of the list
	if l.tail != nil {
		l.tail.next = newNode
	} else {
		l.head = newNode
	}
	l.tail = newNode
}

// Shift remove value from the front of the list. Same as pop in the reversed linked-list
func (l *List) Shift() (interface{}, error) {
	l.Reverse()
	v, err := l.Pop()
	l.Reverse()
	return v, err
}

// Pop remove value from the back of the list
func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("list is empty")
	}
	removedValue := l.tail.Value
	l.tail = l.tail.Prev()
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = l.tail
	}
	return removedValue, nil
}

// Reverse the linked list
func (l *List) Reverse() {
	for node := l.head; node != nil; node = node.Prev() {
		node.next, node.prev = node.prev, node.next
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
