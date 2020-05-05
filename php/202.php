<?php
#!/usr/bin/env php


class Solution
{

    /**
     * HashMap
     */

    // public $hash = [];

    /**
     * @param Integer $n
     * @return Boolean
     */
    // function isHappy($n)
    // {
    //     if ($n == 1) return true;
    //     $count = strlen($n);
    //     $num = 0;
    //     $n = (string) $n;
    //     for ($i = 0; $i < $count; $i++) {
    //         $num += $n[$i] * $n[$i];
    //     }

    //     if (isset($this->hash[$num])) {
    //         return false;
    //     } else {
    //         $this->hash[$num] = 1;
    //     }
    //     return $num == 1 ? true : $this->isHappy($num);
    // }

    /* END */

    /**
     * fast and slow head
     */

    /**
     * @param Integer $n
     * @return Boolean
     */
    function isHappy($n)
    {
        $slow = $n;
        $fast = $this->step($n);
        while ($fast != 1 && $fast != $slow) {
            $slow = $this->step($slow);
            $fast = $this->step($this->step($fast));
        }
        return $fast == 1;
    }

    function step($n)
    {
        $sum = 0;
        $count = strlen($n);
        $n = (string) $n;
        for ($i = 0; $i < $count; $i++) {
            $sum += $n[$i] * $n[$i];
        }
        return $sum;
    }
}

/** TESTING CODE**/
$solutionObj = new Solution();
$result = $solutionObj->isHappy(100);
var_dump($result);
/** END **/
