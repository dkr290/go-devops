package main

import "fmt"

type SingleLinkList struct {
	head *SLLNode
	tail *SLLNode
}

func NewSingleLinkedList() *SingleLinkList {
	return new(SingleLinkList)
}

func (list *SingleLinkList) Add(v int) {
	newNode := &SLLNode{value: v}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}

	list.tail = newNode

}

func (list *SingleLinkList) String() string {
	s := ""
	for n := list.head; n != nil; n = n.next {
		s += fmt.Sprintf(" {%d}", n.GetValue())
	}
	return s
}
