package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
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

func ListReverse(l *List) interface{} {
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
