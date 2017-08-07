package main

import (
	tph "github.com/NerdDiffer/toyproblems/helpers"
)

var joinedList = new(tph.SList)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *tph.ListNode, l2 *tph.ListNode) *tph.ListNode {
	if l1 == nil && l2 == nil { // are both lists done?
		return joinedList.Head
	} else if l1 == nil && l2 != nil {
		joinedList.Push(l2.Val)
		return mergeTwoLists(l1, l2.Next())
	} else if l1 != nil && l2 == nil {
		joinedList.Push(l1.Val)
		return mergeTwoLists(l1.Next(), l2)
	} else {
		// take the lower value
		if l1.Val <= l2.Val {
			joinedList.Push(l1.Val)
			return mergeTwoLists(l1.Next(), l2)
		} else {
			joinedList.Push(l2.Val)
			return mergeTwoLists(l1, l2.Next())
		}
	}
}

func main() {
	l1 := new(tph.SList)
	l2 := new(tph.SList)

	l1.Push(10)
	l1.Push(30)
	l1.Push(50)
	l2.Push(20)
	l2.Push(40)

	l1.Print()
	l2.Print()
	mergeTwoLists(l1.Head, l2.Head)
	joinedList.Print()
}
