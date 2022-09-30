package main

import (
	"fmt"
)

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "three")
	ListPushBack(link, 3)
	ListPushBack(link, "1")

	fmt.Println(ListLast(link))
	fmt.Println(ListLast(link2))
}

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
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

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	} else {
		current := l.Head
		cmpt := -1
		for current.Next != nil {
			current = current.Next
			cmpt++
		}
		return cmpt
	}
}
