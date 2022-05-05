package solution

import "go_leetcode/common"

type ListNode = common.ListNode

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = nil
	var trav *ListNode = nil
	var carry int = 0
	for l1 != nil || l2 != nil || carry > 0 {
		newNode := new(ListNode)
		if (l1 == nil && l2 == nil) {
			newNode.Val = carry
			carry = 0
		} else if (l1 == nil) {
			newNode.Val = (l2.Val + carry)
			carry = newNode.Val / 10
			newNode.Val %= 10
			l2 = l2.Next
		} else if (l2 == nil) {
			newNode.Val = (l1.Val + carry)
			carry = newNode.Val / 10
			newNode.Val %= 10
			l1 = l1.Next
		} else {
			newNode.Val = l1.Val + l2.Val + carry
			carry = newNode.Val / 10
			newNode.Val %= 10
			l1 = l1.Next
			l2 = l2.Next
		}
		if (result == nil) {
			result = newNode
			trav = result
		} else {
			trav.Next = newNode
			trav = trav.Next
		}
	}
	return result
}
