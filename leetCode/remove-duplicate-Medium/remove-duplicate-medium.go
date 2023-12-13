package main

import "fmt"

func removeDuplicates(nums []int) int {
	var count int
	var limit int
	varLimit := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] != varLimit {
			limit = 0
			varLimit = nums[i]
		}

		if limit < 2 {
			nums[count] = nums[i]
			count++
			limit++
		}
	}

	return count
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates(nums))
}
