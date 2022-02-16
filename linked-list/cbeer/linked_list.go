package linkedlist

import (
	"errors"
)

var ErrEmptyList = errors.New("empty list")

// Node holds a value and pointers to the next and previous node.
type Node struct {
	previousNode *Node
	nextNode     *Node
	Val          interface{}
}

// List holds references to the first and last node.
type List struct {
	firstNode *Node
	lastNode  *Node
}

// Given an array of values, return a doubly linked list of values.
func NewList(args ...interface{}) *List {
	var newList List
	for _, value := range args {
		newList.PushBack(value)
	}
	return &newList
}

// Return the next node in the linked list
func (n *Node) Next() *Node {
	return n.nextNode
}

// Return the previous node in the linked list
func (n *Node) Prev() *Node {
	return n.previousNode
}

// insert the value at the front of the list
func (l *List) PushFront(v interface{}) {
	newNode := &Node{Val: v, previousNode: nil, nextNode: nil}
	if l.lastNode == nil {
		l.lastNode = newNode
	} else {
		prevFirstNode := l.firstNode
		newNode.nextNode = prevFirstNode
		prevFirstNode.previousNode = newNode
	}
	l.firstNode = newNode
}

// insert the value at the back of the list
func (l *List) PushBack(v interface{}) {
	newNode := &Node{Val: v, previousNode: nil, nextNode: nil}
	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		prevLastNode := l.lastNode
		newNode.previousNode = prevLastNode
		prevLastNode.nextNode = newNode
	}
	l.lastNode = newNode
}

// remove the value from the front of the list and return it
func (l *List) PopFront() (interface{}, error) {
	if l.firstNode == nil {
		return nil, ErrEmptyList
	}
	frontNode := l.First()
	if frontNode.nextNode != nil {
		l.firstNode = frontNode.nextNode
		frontNode.nextNode.previousNode = nil
	} else {
		l.firstNode = nil
		l.lastNode = nil
	}
	return frontNode.Val, nil
}

// remove the value from the back of the list and return it
func (l *List) PopBack() (interface{}, error) {
	if l.lastNode == nil {
		return nil, ErrEmptyList
	}
	backNode := l.Last()
	if backNode.previousNode != nil {
		l.lastNode = backNode.previousNode
		backNode.previousNode.nextNode = nil
	} else {
		l.firstNode = nil
		l.lastNode = nil
	}
	return backNode.Val, nil
}

// reverse the linked list
func (l *List) Reverse() {
	if l.firstNode != l.lastNode {
		// This is not a great way to do this.
		// At a minimum I should use an array and NewList.
		// Best option would be to switch the list inplace
		// But I wanted to play around with Push and Pop and I'm too lazy to fix it now.
		var tempList List
		newValue, err := l.PopBack()
		for err == nil {
			tempList.PushFront(newValue)
			newValue, err = l.PopBack()
		}
		newValue, err = tempList.PopBack()
		for err == nil {
			l.PushBack(newValue)
			newValue, err = tempList.PopBack()
		}
	}
}

// returns a pointer to the first node in the list
func (l *List) First() *Node {
	return l.firstNode
}

// returns a pointer to the last node in the list
func (l *List) Last() *Node {
	return l.lastNode
}
