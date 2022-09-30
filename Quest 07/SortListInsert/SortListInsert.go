package main

import (
	"fmt"
)

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}

type NodeI struct {
	Data interface{}
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	var current *NodeI
	var current2 *NodeI
	var temp int

	current = l
	for current.Next != nil {
		current2 = current.Next
		for current2 != nil {
			if current.Data.(int) > current2.Data.(int) {
				temp = current.Data.(int)
				current.Data = current2.Data
				current2.Data = temp
			}
			current2 = current2.Next
		}
		current = current.Next
	}
	return l
}
