package main

import "fmt"

func removeElement(nums []int, val int) int {

	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[start] = nums[i]
			start++
		}
	}
	return start
}

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
