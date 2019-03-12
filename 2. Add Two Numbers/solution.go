package ___Add_Two_Numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1 := l1
	p2 := l2
	res := &ListNode{Val: 0}
	p := res
	carry := 0
	for p1 != nil && p2 != nil {
		sum := p1.Val + p2.Val + carry
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		carry = sum / 10
		p1 = p1.Next
		p2 = p2.Next
	}

	for p1 != nil {
		sum := p1.Val + carry
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		carry = sum / 10
		p1 = p1.Next
	}

	for p2 != nil {
		sum := p2.Val + carry
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
		carry = sum / 10
		p2 = p2.Next
	}

	if carry == 1 {
		p.Next = &ListNode{Val: 1}
	}

	return res.Next
}
