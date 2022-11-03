package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	//     s, f := head, head

	//     for f != nil && f.Next != nil {
	//         s = s.Next
	//         f = f.Next.Next
	//     }

	//     return s

	result := []*ListNode{}

	for head != nil {
		result = append(result, head)

		head = head.Next
	}

	return result[len(result)/2]
}
