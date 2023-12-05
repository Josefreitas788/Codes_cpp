package main

import "fmt"

func removeDuplicates(nums []int) (int, []int) {
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[count] != nums[i] {
			count++
			nums[count] = nums[i]

		}
	}
	return count + 1, nums
}

func main() {
	nums := []int{1, 1, 2}
	n, nums := removeDuplicates(nums)
	fmt.Println(n, nums)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n, nums = removeDuplicates(nums)
	fmt.Println(n, nums)
}
