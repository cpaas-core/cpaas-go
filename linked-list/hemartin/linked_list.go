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
	Nodes []*Node
}

func NewList(args ...interface{}) *List {
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

	return &List{Nodes: nodes}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) PushFront(v interface{}) {
	l.Reverse()
	l.PushBack(v)
	l.Reverse()
}

func (l *List) PushBack(v interface{}) {
	node := &Node{Val: v}
	last := l.Last()
	if last != nil {
		last.next = node
		node.prev = last
	}

	l.Nodes = append(l.Nodes, node)
}

func (l *List) PopFront() (interface{}, error) {
	if len(l.Nodes) == 0 {
		return nil, ErrEmptyList
	}

	var nodes []*Node
	first := l.First()
	for idx, node := range l.Nodes {
		if idx == 0 {
			continue
		}
		if idx == 1 {
			node.prev = nil
		}

		nodes = append(nodes, node)
	}
	l.Nodes = nodes

	return first.Val, nil
}

func (l *List) PopBack() (interface{}, error) {
	l.Reverse()
	ret, err := l.PopFront()
	l.Reverse()
	return ret, err
}

func (l *List) Reverse() {
	for _, node := range l.Nodes {
		node.prev, node.next = node.next, node.prev
	}

	for idx := len(l.Nodes)/2 - 1; idx >= 0; idx-- {
		opposite := len(l.Nodes) - 1 - idx
		l.Nodes[idx], l.Nodes[opposite] = l.Nodes[opposite], l.Nodes[idx]
	}
}

func (l *List) First() *Node {
	if len(l.Nodes) > 0 {
		return l.Nodes[0]
	}
	return nil
}

func (l *List) Last() *Node {
	if len(l.Nodes) > 0 {
		if len(l.Nodes) == 1 {
			return l.Nodes[0]
		}
		return l.Nodes[len(l.Nodes)-1]
	}
	return nil
}

func (l *List) Len() int {
	return len(l.Nodes)
}
