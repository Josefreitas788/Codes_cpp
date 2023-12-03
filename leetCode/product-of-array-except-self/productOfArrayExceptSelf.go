package main

import "fmt"

func productExceptSelf(nums []int) []int {
	product := 1
	productLeft := make([]int, len(nums))
	productRight := make([]int, len(nums))

	productLeft[0] = 1
	for i := 1; i < len(nums); i++ {
		product = nums[i-1] * product
		productLeft[i] = product
	}

	product = 1
	productRight[len(nums)-1] = 1
	for i := len(nums) - 1; i > 0; i-- {
		product = nums[i] * product
		productLeft[i-1] = product
	}

	for i := 0; i < len(nums); i++ {
		productLeft[i] = productLeft[i] * productRight[i]
	}
	return productLeft
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(productExceptSelf(nums))

}
