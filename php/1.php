#!/usr/bin/env php
<?php
class Solution
{

    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum(array $nums, int $target)
    {
        $map = [];
        foreach ($nums as $index => $num) {
            $diff = $target - $num;
            if (isset($map[$diff])) {
                return [$map[$diff], $index];
            }
            $map[$num] = $index;
        }
        return [];
    }
}
$nums = [2, 7, 11, 15]; 
$target = 9;

$solutionObj = new Solution();
$solutionObj->twoSum($nums,$target);