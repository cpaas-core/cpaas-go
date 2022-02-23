package linkedlist

import "errors"

var ErrEmptyList = errors.New("Empty List")

// Define List and Node types here.
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
	newList := &List{}
	if len(args) < 1 {
		return newList
	}
	for _, arg := range args {
		newList.PushBack(arg)
	}
	return newList
}

func (n *Node) Next() *Node {
	if n != nil {
		return n.next
	}
	return nil
}

func (n *Node) Prev() *Node {
	if n != nil {
		return n.prev
	}
	return nil
}

func (l *List) PushFront(v interface{}) {
	if l != nil {
		if l.head == nil {
			l.head = &Node{Val: v}
			l.tail = l.head
		} else {
			prevHead := l.head
			l.head = &Node{prev: nil, Val: v, next: prevHead}
			prevHead.prev = l.head
		}
	}
}

func (l *List) PushBack(v interface{}) {
	if l != nil {
		if l.head == nil {
			l.head = &Node{Val: v}
			l.tail = l.head
		} else {
			prevTail := l.tail
			l.tail = &Node{prev: prevTail, Val: v}
			prevTail.next = l.tail
		}
	}
}

func (l *List) PopFront() (interface{}, error) {
	if l == nil || l.head == nil {
		return nil, ErrEmptyList
	}
	prevHeadValue := l.head.Val
	nextHead := l.head.next
	if nextHead != nil {
		nextHead.prev = nil
		l.head = nextHead
		return prevHeadValue, nil
	}
	l.head = nil
	l.tail = nil
	return prevHeadValue, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l == nil || l.tail == nil {
		return nil, ErrEmptyList
	}
	prevTailValue := l.tail.Val
	prevTail := l.tail.prev
	if prevTail != nil {
		prevTail.next = nil
		l.tail = prevTail
		return prevTailValue, nil
	}
	l.head = nil
	l.tail = nil
	return prevTailValue, nil
}

func (l *List) Reverse() {
	if l == nil || l.head == nil {
		return
	}
	l.tail = l.head
	node := l.head
	var tmp *Node
	for node != nil {
		tmp = node.prev
		node.prev, node.next = node.next, node.prev
		node = node.prev
	}
	if tmp != nil {
		l.head = tmp.prev
	}
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
