package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendOneList(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	l2.Next = l1
	return appendOneList(l1.Next, l2.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sum, rest int
	sumList := &ListNode{}
	head := sumList
	for {
		sum = l1.Val + l2.Val + rest
		if sum >= 10 {
			sum -= 10
			rest = 1
		} else {
			rest = 0
		}
		sumList.Val = sum

		if l1.Next == nil && l2.Next == nil {
			if rest == 1 {
				sumList.Next = &ListNode{}
				sumList = sumList.Next
				sumList.Val = rest
			} else {
				sumList = nil
			}
			break
		}
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
		sumList.Next = &ListNode{}
		sumList = sumList.Next
	}
	return head
}

func main() {
	l1Array := []int{9, 9, 9, 9, 9, 9, 9}
	l1 := &ListNode{}
	l1Head := l1
	for i := 0; i < len(l1Array); i++ {
		l1.Val = l1Array[i]
		l1.Next = &ListNode{}
		l1 = l1.Next
	}
	l1 = nil

	l2Array := []int{9, 9, 9, 9}
	l2 := &ListNode{}
	l2Head := l2
	for i := 0; i < len(l2Array); i++ {
		l2.Val = l2Array[i]
		l2.Next = &ListNode{}
		l2 = l2.Next
	}
	l2 = nil

	l := addTwoNumbers(l1Head, l2Head)
	for {
		fmt.Println(l.Val)
		if l.Next == nil {
			break
		}
		l = l.Next
	}
}
