package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i, j := 0, 0
	mergedArray := make([]int, m+n)
	for i < m || j < n {
		if i == m {
			for j < n {
				mergedArray[i+j] = nums2[j]
				j++
			}
			break
		}
		if j == n {
			for i < m {
				mergedArray[i+j] = nums1[i]
				i++
			}
			break
		}
		if nums1[i] < nums2[j] {
			mergedArray[i+j] = nums1[i]
			i++
		} else {
			mergedArray[i+j] = nums2[j]
			j++
		}
	}

	for i := 0; i < m+n; i++ {
		nums1[i] = mergedArray[i]
	}
	return mergedArray
}

func main() {

	fmt.Println(merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3))

}
