package linkedlist

import (
	"errors"
	"fmt"
)

type IList interface {
	First() *Node
	Last() *Node
	PushBack(v interface{})
	PopBack() (interface{}, error)
	PushFront(v interface{})
	PopFront() (interface{}, error)
	Reverse()
}

type INode interface {
	Val() interface{}
	Next() *Node
	Prev() *Node
}

type List struct {
	first *Node
	last  *Node
}

type Node struct {
	Val  interface{}
	next *Node
	prev *Node
}

var ErrEmptyList = errors.New("Error: empty list")

func (l *List) Print() {
	node := l.First()
	fmt.Print("List{ ")
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next()
	}
	fmt.Println(" }")
}

func NewList(args ...interface{}) *List {
	var l IList
	var n *Node
	var first *Node
	var prev *Node
	for _, v := range args {
		n = &Node{v, nil, prev}
		if first == nil {
			first = n
		}
		if prev != nil {
			prev.next = n
		}
		prev = n
	}
	l = &List{first, prev}
	ret := l.(*List)
	return ret
}

func (n *Node) Next() *Node {
	if n == nil {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n == nil {
		return nil
	}
	return n.prev
}

func (l *List) Clear() {
	l.first = nil
	l.last = nil
}

func (l *List) isEmpty() bool {
	return l.first == nil && l.last == nil
}

func (l *List) PushFront(v interface{}) {
	if v != nil {
		node := &Node{v, nil, nil}
		oldFirst := l.First()
		if oldFirst != nil {
			node.next = oldFirst
			oldFirst.prev = node
		}
		l.first = node
		if l.last == nil {
			l.last = node
		}
	}
}

func (l *List) PushBack(v interface{}) {
	node := &Node{v, nil, nil}
	lastPrev := l.Last()
	if lastPrev != nil {
		node.prev = lastPrev
		lastPrev.next = node
	}

	l.last = node
	if l.first == nil {
		l.first = node
	}
}

func (l *List) PopFront() (interface{}, error) {

	if l.isEmpty() {
		return nil, ErrEmptyList
	}
	v := l.First()
	if v == l.Last() {
		l.Clear()
	}
	if v.Next() != nil {
		l.first = v.Next()
		l.first.prev = nil
	}
	v.next = nil
	return v.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.isEmpty() {
		return nil, ErrEmptyList
	}
	v := l.Last()
	if v == nil {
		return nil, ErrEmptyList
	}
	if v == l.First() {
		l.Clear()
	}
	if v.prev != nil {
		l.last = v.prev
		l.last.next = nil
	}
	v.prev = nil
	return v.Val, nil
}

func (l *List) Reverse() {
	if !l.isEmpty() {

		node := l.Last()

		if node != nil {
			n := node
			var lastNode *Node
			for n.Prev() != nil {
				v := n.Prev()
				n.prev = lastNode
				n.next = v
				lastNode = n
				n = v
			}
			node.prev = nil
			l.first = node
			n.next = nil
			n.prev = lastNode
			l.last = n
		}
	}
}

func (l *List) First() *Node {
	return l.first
}

func (l *List) Last() *Node {
	return l.last
}
