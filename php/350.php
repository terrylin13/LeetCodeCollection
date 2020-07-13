<?php
class Solution
{

    /**
     * @param Integer[] $nums1
     * @param Integer[] $nums2
     * @return Integer[]
     */
    function intersect($nums1, $nums2)
    {
        $ans = [];
        $n1 = count($nums1);
        $n2 = count($nums2);
        if ($n1 == 0 || $n2 == 0) return $ans;

        if ($n1 > $n2) {
            $tmp = $nums1;
            $nums1 = $nums2;
            $nums2 = $tmp;
        }

        $hash = [];
        $result = [];
        foreach ($nums1 as $num) {
            if (!isset($hash[$num])) {
                $hash[$num] = 1;
            } else {
                $hash[$num]++;
            }
        }

        foreach ($nums2 as $num) {
            if (isset($hash[$num]) && $hash[$num] > 0) {
                $result[] = $num;
                $hash[$num]--;
            }
        }

        return $result;
    }
}

var_dump((new Solution())->intersect([1, 23, 4], [1, 2, 4, 55]));
