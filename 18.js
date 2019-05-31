/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[][]}
 */
var fourSum = function(nums, target) {
  nums = nums.sort((a, b) => a - b)
  let i, j, left, ç, sum
  let len = nums.length
  let res = []
  for (i = 0; i < len - 3; i++) {
      if (i > 0 && nums[i] === nums[i - 1]) continue
      for (j = i + 1; j < len - 2; j++) {
          if (j !== i + 1 && nums[j] === nums[j - 1]) continue
          left = j + 1
          right = len - 1
          while (left < right) {
              sum = nums[i] + nums[j] + nums[left] + nums[right]
              if (sum === target) {
                  res.push([nums[i], nums[j], nums[left], nums[right]])
                  while (nums[left + 1] === nums[left]) left++
                  while (nums[right - 1] === nums[right]) right--
                  left++
                  right--
              } else if (sum > target) {
                right--
              } else {
                left++
              }
          }
      }
  }
  return res
};