package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	if lists == nil {
		return nil
	}

	length := len(lists)
	if length == 1 {
		return lists[0]
	}

	if length == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	k := length >> 1
	left := mergeKLists(lists[:k])
	right := mergeKLists(lists[k:length])

	return mergeTwoLists(left, right)

}

func mergeTwoLists(listA *ListNode, listB *ListNode) *ListNode {

	var head, currNode *ListNode

	if listA == nil {
		return listB
	}

	if listB == nil {
		return listA
	}

	if listA.Val > listB.Val {
		head, currNode, listB = listB, listB, listB.Next
	} else {
		head, currNode, listA = listA, listA, listA.Next
	}

	for listA != nil && listB != nil {
		if listA.Val > listB.Val {
			currNode.Next, listB = listB, listB.Next
		} else {
			currNode.Next, listA = listA, listA.Next
		}
		currNode = currNode.Next
	}

	if listA != nil {
		currNode.Next = listA
	} else if listB != nil {
		currNode.Next = listB

	}

	return head

}

func main() {
	node3 := &ListNode{Val: 5, Next: nil}
	node2 := &ListNode{Val: 4, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	node6 := &ListNode{Val: 4, Next: nil}
	node5 := &ListNode{Val: 3, Next: node6}
	node4 := &ListNode{Val: 1, Next: node5}

	node8 := &ListNode{Val: 6, Next: nil}
	node7 := &ListNode{Val: 2, Next: node8}
	lists := []*ListNode{node1, node4, node7}

	list := mergeKLists(lists)
	for list != nil {
		fmt.Print(list.Val, "->")
		list = list.Next
	}

}
