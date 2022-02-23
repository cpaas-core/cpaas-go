package linkedlist

import (
	"errors"
)

var (
	ErrEmptyList = errors.New("empty list")
)

type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

func NewNode(v interface{}) *Node {
	return &Node{Val: v, prev: nil, next: nil}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	first *Node
	last  *Node
}

func NewList(args ...interface{}) *List {
	var ret List
	for _, v := range args {
		ret.PushBack(v)
	}
	return &ret
}

func (l *List) AddFirstNode(node *Node) {
	l.first = node
	l.last = node
}

func (l *List) EmptyList() {
	l.first = nil
	l.last = nil
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}

func (l *List) IsEmpty() bool {
	return l.First() == nil && l.Last() == nil
}

func (l *List) IsNotEmpty() bool {
	return !l.IsEmpty()
}

func (l *List) Len() (len int) {
	if l.IsNotEmpty() {
		for node := l.First(); node != nil; node = node.Next() {
			len++
		}
	}
	return
}

func (l *List) PushFront(v interface{}) {
	node := NewNode(v)
	if l.IsEmpty() {
		l.AddFirstNode(node)
	} else {
		l.first.prev = node
		node.next = l.first
		l.first = node
	}
}

func (l *List) PushBack(v interface{}) {
	node := NewNode(v)
	if l.IsEmpty() {
		l.AddFirstNode(node)
	} else {
		l.last.next = node
		node.prev = l.last
		l.last = node
	}
}

func (l *List) PopFront() (interface{}, error) {
	if l.IsNotEmpty() {
		ret := l.First().Val
		if l.Len() == 1 {
			l.EmptyList()
		} else {
			l.first.next.prev = nil
			l.first = l.First().Next()
		}
		return ret, nil
	}
	return nil, ErrEmptyList
}

func (l *List) PopBack() (interface{}, error) {
	if l.IsNotEmpty() {
		ret := l.Last().Val
		if l.Len() == 1 {
			l.EmptyList()
		} else {
			l.last.prev.next = nil
			l.last = l.Last().Prev()
		}
		return ret, nil
	}
	return nil, ErrEmptyList
}

func (l *List) Reverse() {
	if l.Len() > 1 {
		for node := l.First(); node != nil; node = node.Prev() {
			node.prev, node.next = node.next, node.prev
		}
		l.first, l.last = l.last, l.first
	}
}
