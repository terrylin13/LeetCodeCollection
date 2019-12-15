/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    let res = []
    nums.sort((a,b) => a - b)
    for (let i = 0; i < nums.length; i++) {
        if (nums[i] === nums[i - 1]) continue
        let target = -nums[i]
        let left = i + 1, right = nums.length - 1
        while (left < right) {
            if (nums[left] + nums[right] === target) {
                res.push([nums[left], -target, nums[right]])
                while (nums[left] === nums[left + 1])
                    left++
                left++
                while (nums[right] === nums[right - 1])
                    right--
                right--
            }
            else if (nums[left] + nums[right] > target) right--
            else left++
        }
        
    }
    return res
};
