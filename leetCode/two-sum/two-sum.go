package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if ((nums[i] + nums[j]) == target) && (i != j) {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	array := []int{1, 2, 4, 5, 6, 8, 9, 0, 4, 2, 4}
	target := 14
	result := twoSum(array, target)
	fmt.Print(result)
}
