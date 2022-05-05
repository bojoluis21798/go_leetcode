package common

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

type List struct {
	Head *ListNode
}

func InsertLastFunc(head **ListNode, data int) {
	newNode := &ListNode{ Val: data, Next: nil}
	temp := head
	for *temp != nil {
		 temp = &(*temp).Next
	}
	*temp = newNode
}

func DisplayFunc(list *ListNode) {
	for trav := list; trav != nil; trav = trav.Next {
		fmt.Printf("%d ", trav.Val)
	}
}

func ArrayToListFunc(arr []int) *ListNode {
	var list *ListNode
	for _, num := range arr {
		InsertLastFunc(&list, num)
	}
	return list
}

func (list *List) InsertLast(data int) {
	InsertLastFunc(&list.Head, data)
}

func (list List) Display() {
	DisplayFunc(list.Head)
}

func (list *List) ArrayToList(arr []int) {
	list.Head = ArrayToListFunc(arr)
}
