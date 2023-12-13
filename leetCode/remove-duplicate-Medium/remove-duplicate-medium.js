/**
 * @param {number[]} nums
 * @return {number}
 */
var removeDuplicates = function (nums) {
    count = 0
    limit = 0
    let varLimite = nums[0]

    for (let i = 0; i < nums.length; i++) {
        if (varLimite != nums[i]) {
            limit = 0
            varLimite = nums[i]
        }
        if (limit < 2) {
            limit++
            nums[count] = nums[i]
            count++
        }
    }

    return count
};

// test case 1
console.log(removeDuplicates([1, 1, 1, 2, 2, 3]));

// test case 2
console.log(removeDuplicates([0, 0, 1, 1, 1, 1, 2, 3, 3]));
//expected output  [0,0,1,1,2,2,3,3,4]
