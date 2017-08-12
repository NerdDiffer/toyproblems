package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			l2 = nil
			continue
		} else if l2 == nil {
			curr.Next = l1
			l1 = nil
			continue
		}

		if l1.Val <= l2.Val {
			curr.Next = l1
			curr, l1 = curr.Next, l1.Next
		} else {
			curr.Next = l2
			curr, l2 = curr.Next, l2.Next
		}
	}

	return result.Next
}

// for quick & dirty QA, change package name & this function name to `main`
func run_mergeTwoLists() {
	l1 := &ListNode{Val: 10}
	l1.Next = &ListNode{Val: 30}
	l1.Next.Next = &ListNode{Val: 50}

	l2 := &ListNode{Val: 20}
	l2.Next = &ListNode{Val: 40}

	printListNode(l1)
	printListNode(l2)

	merged := mergeTwoLists(l1, l2)
	printListNode(merged)
}

func printListNode(n *ListNode) {
	curr := n

	for curr != nil {
		fmt.Printf("%+v -> ", curr.Val)
		curr = curr.Next
	}

	fmt.Println()
}
