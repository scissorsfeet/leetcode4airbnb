package ___Add_Two_Numbers

import (
	"testing"
)

func l2int(p *ListNode) int {
	unit := 1
	res := 0
	for p != nil {
		res += p.Val * unit
		unit *= 10
		p = p.Next
	}
	return res
}

func TestA(t *testing.T) {
	l1 := &ListNode{Val: 1}

	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}

	res := addTwoNumbers(l1, l2)
	if l2int(res) != 1000 {
		t.Fatal(l2int(res))
	}
}

func TestB(t *testing.T) {
	var l1 *ListNode = nil

	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}

	res := addTwoNumbers(l1, l2)
	if l2int(res) != 999 {
		t.Fatal(l2int(res))
	}
}
