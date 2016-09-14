package datastructure

import (
	"testing"
)
/*
func Test_NewSinglyList(*testing.T)  {
	node := NewSinglyListByValue(1)
	Show(node)

	node = NewSinglyListByValue(109.1)
	Show(node)

	node = NewSinglyListByValue("Hello World")
	Show(node)


	list := []int{1, 2, 3}
	node = NewSinglyListByValue(list)
	Show(node)
}

func Test_SinglyListByArrayFromHead(*testing.T) {
	num1 := []int{2, 4, 3}
	num2 := []int{5, 6, 4}

	n1 := SinglyListByArrayFromHead(num1)
	n2 := SinglyListByArrayFromHead(num2)

	Show(n1)
	Show(n2)
}

func Test_SinglyListByArrayFromTail(*testing.T) {
	num1 := []int{2, 4, 3}
	num2 := []int{5, 6, 4}

	n1 := SinglyListByArrayFromTail(num1)
	n2 := SinglyListByArrayFromTail(num2)

	Show(n1)
	Show(n2)
}

*/
func Test_Add(*testing.T) {
	//num1 := []int{2, 4, 3}
	//num2 := []int{5, 6, 4}

	//num1 := []int{5}
	//num2 := []int{5}

	//num1 := []int{9, 8}
	//num2 := []int{1}

	//num1 := []int{1}
	//num2 := []int{9, 9, 9}

	n1 := SinglyListByArrayFromTail(num1)
	n2 := SinglyListByArrayFromTail(num2)

	Show(n1)
	Show(n2)

	n3 := Add(n1, n2)
	Show(n3)
}
