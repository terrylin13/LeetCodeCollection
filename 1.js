/**
 * @file Two Sum Solution
 * @link https://leetcode.com/problems/two-sum/
 */

/**
 * Brute force
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    let length = nums.length
    for(let i=0;i<length;i++){
      for(let j=i+1;j<length;j++){
          if(nums[i]+nums[j]==target){
           console.log(i,j)
           return [i,j];
          }
      }
    }
    console.log('none')
    return [];
};


/**
 * Difference mapping table 
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
  var i,j,len = nums.length
  var hash = {}
  for(i=0;i<len;i++){
    j = target - nums[i]
    if(hash[nums[i]] !== undefined){
      console.log(i,hash[nums[i]])
      return [hash[nums[i]],i]
    }else{
      hash[j] = i;
    }
  }
};


twoSum([2,7,11,15],9)