var merge = function(nums1, m, nums2, n) {
	i = 0
	j = 0

	mergedArray = []

	while (i < m || j < n) {
		if (i == m) {
			mergedArray = mergedArray.concat(nums2.slice(j))
			break
		} else if (j == n) {
			mergedArray = mergedArray.concat(nums1.slice(i,m))
			break
		} else
		if (nums1[i] > nums2[j]) {
			mergedArray.push(nums2[j])
			j++		
		} else{
			mergedArray.push(nums1[i])
			i++
			
		}
	}

	return mergedArray

};

// test case 0
console.log(merge([2,0], 1, [1], 1))

// test case 1
console.log(merge([1,2,3,0,0,0], 3, [2,5,6], 3))

// test case 2
console.log(merge([1], 1, [], 0))

// test case 3
console.log(merge([0], 0, [1], 1))