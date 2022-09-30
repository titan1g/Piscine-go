package main

import (
	"fmt"
)

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "hello1")
	ListPushBack(link, "hello2")
	ListPushBack(link, "hello3")

	found := ListFind(link, interface{}("hello2"), CompStr)

	fmt.Println(found)
	fmt.Println(*found)
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

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	iterator := l.Head
	for iterator != nil {
		if comp(iterator.Data, ref) {
			return &iterator.Data
		}
		iterator = iterator.Next
	}
	return nil

}

func CompStr(a, b interface{}) bool {
	if a.(string) == b.(string) {
		return true
	}
	return false
}
