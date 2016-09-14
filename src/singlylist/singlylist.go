package datastructure

import (
	"fmt"
)

type SinglyList struct {
	//Value interface{}
	Value int
	Next  *SinglyList
}

func NewSinglyList() *SinglyList {
	node := new(SinglyList)
	node.Value = 1
	node.Next = nil
	return node
}

func NewSinglyListByValue(value interface{}) *SinglyList {
	node := new(SinglyList)
	//node.Value = value
	node.Next = nil
	return node
}

//func SinglyListByArrayFromHead(array interface{}) *SinglyList {
func SinglyListByArrayFromHead(array []int) *SinglyList {
	head := new(SinglyList)
	head.Next = nil

	for _, value := range array {
		node := new(SinglyList)
		node.Value = value
		node.Next = head.Next
		head.Next = node
	}

	return head.Next
}

//func SinglyListByArrayFromTail(array interface{}) *SinglyList {
func SinglyListByArrayFromTail(array []int) *SinglyList {
	head := new(SinglyList)
	head.Next = nil

	tail := head

	for _, value := range array {
		node := new(SinglyList)
		node.Value = value
		tail.Next = node
		tail = node
	}

	return head.Next
}

/*
func Add(a, b *SinglyList) *SinglyList {
	var over int
	head := new(SinglyList)
	head.Next = nil
	tail := head

	for {

		node := new(SinglyList)
		node.Value = a.Value + b.Value + over
		fmt.Println(node.Value)
		if node.Value >= 10 {
			node.Value = node.Value % 10
			over = 1
		} else {
			over = 0
		}

		tail.Next = node
		tail = node

		a = a.Next
		b = b.Next

		if a == nil || b == nil {
			break
		}

	}

	for {

		if a != nil {
			node := new(SinglyList)
			node.Value = a.Value + over
			tail.Next = node
			tail = node

			a = a.Next
		} else {
			break
		}

	}

	for {

		if b != nil {
			node := new(SinglyList)
			node.Value = b.Value + over
			tail.Next = node
			tail = node

			b = b.Next
		} else  {
			break
		}

	}

	if over == 1 {
		node := new(SinglyList)
		node.Value = over
		tail.Next = node
		tail = node
	}

	return head.Next
}
*/

func Sum(v1, v2, ov int) (sum, over int) {
	sum = v1 + v2 + ov
	if sum >= 10 {
		sum = sum % 10
		over = 1
		return sum, over
	} else {
		return sum, 0
	}

}

func Add(l1, l2 *SinglyList) *SinglyList {
	head := new(SinglyList)
	head.Next = nil

	tail := head

	var over int

	for {
		node := new(SinglyList)
		node.Value, over = Sum(l1.Value, l2.Value, over)

		tail.Next = node
		tail = node

		l1 = l1.Next
		l2 = l2.Next

		/*
		if l1 == nil && l2 == nil && over == 1 {
			node := new(SinglyList)
			node.Value = over

			tail.Next = node
			tail = node
		}
		*/

		if l1 == nil || l2 == nil {
			break
		}
	}

	for {

		if l1 != nil {
			node := new(SinglyList)
			node.Value, over = Sum(l1.Value, 0, over)
			tail.Next = node
			tail = node

			l1 = l1.Next
		} else {
			break
		}

	}

	for {

		if l2 != nil {
			node := new(SinglyList)
			node.Value, over = Sum(l2.Value, 0, over)
			tail.Next = node
			tail = node

			l2 = l2.Next
		}  else {
			break
		}

	}

	if over == 1 {
		node := new(SinglyList)
		node.Value = 1
		tail.Next = node
		tail = node
	}

	return head.Next
}

func Show(s *SinglyList) {
	p := s

	fmt.Println("---------------------")
	for {
		fmt.Println(p.Value)
		p = p.Next
		if p == nil {
			fmt.Println("---------------------")
			return
		}
	}
}
