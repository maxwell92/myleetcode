package main

import (
	"fmt"
)

func isExisted(nums []int, other, current int) int {
	for index, value := range nums {
		if value == other && index != current{
			return index	
		}	
	}
	return -1
}
/*
func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	sort.Ints(nums)
	fmt.Printf("%v", nums)
	for index, value := range nums {
		other := target - value	
		if otherindex := isExisted(nums, other); otherindex != -1 {
			result = append(result, index)	
			result = append(result, otherindex)
			return result
		}
	}				

	return nil
}
*/

func twoSum(nums []int, target int) []int {
	result := make([]int, 0)
	for index, value := range nums {
		other := target - value
		if otherindex := isExisted(nums, other, index); otherindex != -1 {
			result = append(result, index)	
			result = append(result, otherindex)
			return result
		}
	}
	return nil
}



func main() {
	target := 6
	//nums := []int{2, 7, 11, 15}
	nums := []int{3, 2, 4}
	index := twoSum(nums, target)
	fmt.Printf("%v", index)
}
