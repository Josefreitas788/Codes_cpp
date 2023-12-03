var productExceptSelf = function(nums) {
    let productLeft = [];
    let productRight = [];
    let product = 1;

    for (let i = 0; i < nums.length; i++) {
        productLeft.push(product);
        product *= nums[i];
    }
    
    product = 1;
    for (let i = nums.length - 1; i >= 0; i--) {
        productRight[i] = product;
        product *= nums[i];
    }

    let result = [];
    for (let i = 0; i < nums.length; i++) {
        result.push(productLeft[i] * productRight[i]);
    }

    return result;

};

console.log(productExceptSelf([-1,1,0,-3,3])); // [24,12,8,6]