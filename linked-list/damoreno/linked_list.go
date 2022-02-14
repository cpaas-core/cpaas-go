package linkedlist

import (
	"errors"
)

type List struct {
	head, tail *Node
}

type Node struct {
	Val        interface{}
	next, prev *Node
}

var ErrEmptyList = errors.New("list is empty")

func NewList(args ...interface{}) *List {
	list := new(List)

	for _, value := range args {
		list.PushBack(value)
	}

	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	if l.head == nil {
		l.head = &Node{
			Val: v,
		}
		l.tail = l.head
	} else {
		l.head.prev = &Node{
			Val:  v,
			next: l.head,
		}
		l.head = l.head.prev
	}
}

func (l *List) PushBack(v interface{}) {
	if l.tail == nil {
		l.tail = &Node{
			Val: v,
		}
		l.head = l.tail
	} else {
		l.tail.next = &Node{
			Val:  v,
			prev: l.tail,
		}
		l.tail = l.tail.next
	}
}

func (l *List) PopFront() (interface{}, error) {
	if l.head == nil {
		return nil, ErrEmptyList
	}

	value := l.head.Val

	if l.head.next == nil {
		l.head, l.tail = nil, nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	return value, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	value := l.tail.Val

	if l.tail.prev == nil {
		l.head, l.tail = nil, nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	return value, nil
}

func (l *List) Reverse() {
	node := l.head

	for node != nil {
		node.prev, node.next = node.next, node.prev
		node = node.prev
	}

	l.head, l.tail = l.tail, l.head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
