package utils

import (
	"fmt"
	"go_leetcode/common"
)

func InsertLast(head **common.ListNode, data int) {
	newNode := &common.ListNode{ Val: data, Next: nil}
	temp := head
	for *temp != nil {
		 temp = &(*temp).Next
	}
	*temp = newNode
}

func Display(list *common.ListNode) {
	for trav := list; trav != nil; trav = trav.Next {
		fmt.Println(trav.Val)
	}
}

func ArrayToList(arr []int) *common.ListNode {
	var list *common.ListNode
	for _, num := range arr {
		InsertLast(&list, num)
	}
	return list
}
