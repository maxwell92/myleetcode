package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
/*
func newNum(num []int) *ListNode {
	head := new(ListNode)
	head.Val = -1
	head.Next = nil
	size := len(num)
	for i := 0; i < size; i++ {
		node := new(ListNode)	
		node.Val = num[i]
		node.Next = head.Next
		head.Next = node

		fmt.Println(node.Val)
	}
	return head.Next
}
*/
func newNum(num []int) *ListNode {
	head := new(ListNode)
	head.Val = -1
	head.Next = nil

	tail := head 

	size := len(num)
	for i := 0; i < size; i++ {
		node := new(ListNode)	
		node.Val = num[i]
		tail.Next = node
		tail = node	
	}
	return head.Next
}

func showNum(num *ListNode) {
	p := num
	for {
		fmt.Println(p.Val)
		p = p.Next
		if p == nil {
			return	
		}	
	}	
}

func main() {
	nums1 := []int{2, 4, 3}
	nums2 := []int{5, 6, 4}

	num1 := newNum(nums1) 
	num2 := newNum(nums2)

	showNum(num1)
	showNum(num2)

}
