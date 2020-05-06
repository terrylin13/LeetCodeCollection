<?php
class Solution
{

    /**
     * @param Integer[] $A
     * @param Integer[] $B
     * @param Integer[] $C
     * @param Integer[] $D
     * @return Integer
     */
    function fourSumCount($A, $B, $C, $D)
    {
        $res = 0;
        $cache = array();
        foreach ($A as $num1) {
            foreach ($B as $num2) {
                $key =  ($num1 + $num2) * -1;
                $cache[$key] = isset($cache[$key]) ? $cache[$key] + 1 : 1;
            }
        }
        foreach ($C as $num3) {
            foreach ($D as $num4) {
                $key =  ($num3 + $num4);
                if (isset($cache[$key])) $res += $cache[$key];
            }
        }
        return $res;
    }
}
