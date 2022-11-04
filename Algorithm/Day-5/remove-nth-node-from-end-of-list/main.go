package removenthnodefromendoflist

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmpNode := head
	pointer := 0

	for tmpNode != nil {
		pointer++
		tmpNode = tmpNode.Next
	}

	rmvPos := pointer - n

	if n != 0 {
		if rmvPos > 0 {
			tmpNode = head

			for i := 0; i < rmvPos-1; i++ {
				tmpNode = tmpNode.Next
			}

			tmpNode.Next = tmpNode.Next.Next
		} else if pointer == 1 && n == 1 {
			head = nil
		} else if pointer == n {
			head = head.Next
		}
	}

	return head
}

func main() {
	//
}
