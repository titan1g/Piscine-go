package main

import (
	"fmt"
)

type NodeI struct {
	Data interface{}
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

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
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeI{Data: data}

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

func SortedListMerge(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var l *NodeI
	var iterator *NodeI

	l = ListSort(l1)
	iterator = l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = ListSort(l2)
	return l
}
