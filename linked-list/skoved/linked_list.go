package linkedlist

import "errors"

// Define List and Node types here.
var ErrEmptyList = errors.New("empty list")

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	var newList List
	for _, arg := range args {
		newList.PushBack(arg)
	}
	return &newList
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

	val := l.head.Val
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}

	return val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrEmptyList
	}

	val := l.tail.Val
	if l.tail.prev == nil {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}

	return val, nil
}

func (l *List) Reverse() {
	for curNode := l.head; curNode != nil; curNode = curNode.prev {
		curNode.prev, curNode.next = curNode.next, curNode.prev
	}

	tmp := l.head
	l.head = l.tail
	l.tail = tmp
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
