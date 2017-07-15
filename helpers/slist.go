package helpers

import "fmt"

type ListNode struct {
	Val  interface{}
	next *ListNode
}

func (node *ListNode) Next() *ListNode {
	return node.next
}

type SList struct {
	Head *ListNode
}

func (l *SList) Push(val interface{}) {
	node := &ListNode{next: nil, Val: val}

	if l.Head == nil {
		l.Head = node
		return
	}

	curr := l.Head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = node
}

func (l *SList) Print() {
	curr := l.Head

	for curr != nil {
		fmt.Printf("%+v -> ", curr.Val)
		curr = curr.next
	}

	fmt.Println()
}
