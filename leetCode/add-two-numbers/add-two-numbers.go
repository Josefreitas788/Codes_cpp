package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumValue int
	decimal := 1
	for {
		sumValue = sumValue + (l1.Val+l2.Val)*decimal

		if l1.Next == nil || l2.Next == nil {
			break
		}
		l1, l2 = l1.Next, l2.Next

	}
	var sumList *ListNode

}

func main() {

}
