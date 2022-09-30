package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) interface{} {
	if l.Head == nil {
		return 0
	} else {
		current := l.Head
		cmpt := 1
		for current.Next != nil {
			current = current.Next
			cmpt++
		}
		return cmpt
	}
}
func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = n
		return
	}

	iterator := l.Head
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
}
