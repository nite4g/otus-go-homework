package hw04_lru_cache

import (
	"fmt"
	"log"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(interface{}) *ListItem
	PushBack(interface{}) *ListItem
	Remove(*ListItem)
	MoveToFront(*ListItem)
	Scan()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	Head, Tail *ListItem
	Length     int
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.Head
}

func (l *list) Back() *ListItem {
	return l.Tail
}

func (l *list) PushFront(value interface{}) *ListItem {
	var node ListItem

	if l.Head == nil {
		node.Prev = nil
		node.Next = nil
		node.Value = value
		l.Head = &node
		l.Tail = l.Head
	} else {
		l.Front().Prev = &node
		node.Next = l.Front()
		node.Prev = nil
		node.Value = value
		l.Head = &node
	}

	l.Length++
	return &node
}

func (l *list) PushBack(value interface{}) *ListItem {
	var node ListItem

	if l.Tail == nil {
		node.Prev = nil
		node.Next = nil
		node.Value = value
		l.Tail = &node
		l.Head = l.Tail
	} else {
		l.Back().Next = &node
		node.Prev = l.Back()
		node.Next = nil
		node.Value = value
		l.Tail = &node
	}

	l.Length++
	return &node
}

func (l *list) Remove(node *ListItem) {
	if node == nil {
		return
	}

	if node.Next != nil && node.Prev != nil {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		l.Length--
		return
	}

	if node.Next == nil && node.Prev != nil {
		l.Tail = node.Prev
		node.Prev.Next = nil
		l.Length--
		return
	}

	if node.Prev == nil && node.Next != nil {
		l.Head = node.Next
		node.Next.Prev = nil
		l.Length--
		return
	}
	if node.Next == nil && node.Prev == nil {
		// queue could not remove itself
		return
	}
	log.Fatalln("Undef Case for queue Remove() method")
}

func (l *list) MoveToFront(node *ListItem) {
	if node == nil {
		// do if node really exists
		return
	}

	if l.Length == 1 {
		// single element. do nothing
		return
	}

	l.Remove(node)
	l.Length++ // bacause Remove() always do length--
	l.Head.Prev = node
	node.Next = l.Head
	node.Prev = nil
	l.Head = node
}

func (l *list) Scan() {
	// For testing purposes
	var root *ListItem
	root = l.Head
	for root != nil {
		fmt.Printf("%#v\t", root.Value)
		root = root.Next
	}
}
