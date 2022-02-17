package linkedlist

import (
	"errors"
)

var ErrEmptyList error = errors.New("Error")

type Node struct {
	Val  interface{}
	prev *Node
	next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func NewList(args ...interface{}) *List {
	if len(args) == 0 {
		return &List{}
	}

	var nodes []*Node
	for _, value := range args {
		nodes = append(nodes, &Node{
			Val: value,
		})
	}

	if len(args) > 1 {
		for idx, node := range nodes {
			if idx == len(nodes)-1 {
				node.prev = nodes[idx-1]
				node.next = nil
			} else if idx == 0 {
				node.prev = nil
				node.next = nodes[idx+1]
			} else {
				node.prev = nodes[idx-1]
				node.next = nodes[idx+1]
			}
		}
	}

	return &List{Head: nodes[0], Tail: nodes[len(nodes)-1]}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	node := &Node{Val: v}
	first := l.First()
	last := l.Last()

	if last == nil {
		l.Tail = node
	} else {
		node.next = first
		first.prev = node
	}

	l.Head = node
}

func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v}
	first := l.First()
	last := l.Last()

	if first == nil {
		l.Head = node
	} else {
		node.prev = last
		last.next = node
	}

	l.Tail = node
}

func (l *List) PopFront() (interface{}, error) {
	if l.Head == nil {
		return nil, ErrEmptyList
	}

	first := l.First()
	if first.next != nil {
		second := first.next
		second.prev = nil
		l.Head = second
	} else {
		l.Head, l.Tail = nil, nil
	}

	return first.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.Tail == nil {
		return nil, ErrEmptyList
	}

	last := l.Last()
	if last.prev != nil {
		beforeLast := last.prev
		beforeLast.next = nil
		l.Tail = beforeLast
	} else {
		l.Head, l.Tail = nil, nil
	}
	return last.Val, nil
}

// Not working
func (l *List) Reverse() {
	// Empty list
	if l.First() == nil || l.Last() == nil {
		return
	}

	// Single item list
	if l.First() == l.Last() {
		return
	}

	node := l.First()
	for node.next != nil {
		node.prev, node.next = node.next, node.prev
		node = node.prev
	}

	l.Head, l.Tail = l.Tail, l.Head
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Len() int {
	count := 0
	node := l.First()
	if node != nil {
		count++
		for node.next != nil {
			count++
			node = node.next
		}
	}

	return count
}
