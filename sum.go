package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

func NewListNode(array []int) *Node {
	var head *Node
	var curl *Node
	head = nil
	for k, v := range array {
		temp := &Node{Data: v, Next: nil}
		if k == 0 {
			head = temp
			curl = temp
		} else {
			curl.Next = temp
			curl = temp
		}
	}
	return head
}

func AddTwoNumber(l1 *Node, l2 *Node) *Node {
	return AddNode(l1, l2, 0)
}

func AddNode(l1 *Node, l2 *Node, carry int) *Node {
	if l1 == nil && l2 == nil {
		if carry != 0 {
			return &Node{Data: carry, Next: nil}
		} else {
			return nil
		}
	}
	if l1 == nil && l2 != nil {
		l1 = &Node{Data: 0, Next: nil}
	}
	if l1 != nil && l2 == nil {
		l2 = &Node{Data: 0, Next: nil}
	}

	sum := l1.Data + l2.Data + carry
	curl := &Node{Data: sum % 10, Next: nil}
	curl.Next = AddNode(l1.Next, l2.Next, sum/10)

	return curl
}

func (l *Node) PrintList() {
	p := l
	for p != nil {
		fmt.Println(p.Data)
		p = p.Next
	}
}

func ReverseList(head *Node) *Node {
	var next *Node
	var pre *Node
	pre = nil
	for head != nil {
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func sortList(list *Node) *Node {
	var sortedTail *Node

	pStart := &Node{}
	pStart.Next = list
	sortedTail = pStart
	for sortedTail.Next != nil {
		minNode := sortedTail.Next
		temp := sortedTail.Next.Next
		for temp != nil {
			if temp.Data < minNode.Data {
				minNode = temp
			}
			temp = temp.Next
		}
		minNode.Data, sortedTail.Next.Data = sortedTail.Next.Data, minNode.Data
		sortedTail = sortedTail.Next
	}

	list = pStart.Next
	return list
}

func main() {
	listNode := NewListNode([]int{3, 4, 5})
	listNode2 := NewListNode([]int{9, 8, 5})

	list := AddTwoNumber(listNode, listNode2)
	list.PrintList()

	fmt.Println("====Re")
	rlist := ReverseList(list)
	rlist.PrintList()

	listNode3 := NewListNode([]int{9, 1, 0, 7, 3, 8, 5})
	slist := sortList(listNode3)
	slist.PrintList()
}
