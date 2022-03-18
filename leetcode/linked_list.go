package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var head *ListNode
	node4 := new(ListNode)
	node4.Val = -4
	node3 := new(ListNode)
	node3.Val = 0
	node3.Next = node4
	node2 := new(ListNode)
	node2.Val = 2
	node2.Next = node3
	head = new(ListNode)
	head.Val = 2
	head.Next = node2
	node4.Next = head
	fmt.Println(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
