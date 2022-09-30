package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
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
	if l == nil {
		return
	}

	var new_node *NodeL = new(NodeL)
	new_node.Data = data

	if l.Head == nil {
		l.Head = new_node
		l.Tail = new_node
		return
	}

	l.Tail.Next = new_node
	l.Tail = new_node
}
func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil {
		return nil
	}
	cmpt := 0
	for l.Next != nil {
		if cmpt == pos {
			return l
		}
		l = l.Next
		cmpt++
	}
	return nil
}
