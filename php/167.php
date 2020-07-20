<?php
class Solution
{

    /* mapping */
    // function twoSum($numbers, $target)
    // {
    //     $map = [];
    //     foreach ($numbers as $index => $num) {
    //         $diff = $target - $num;
    //         if (isset($map[$diff])) {
    //             return [$map[$diff], $index + 1];
    //         }
    //         $map[$num] = $index + 1;
    //     }
    //     return [];
    // }


    /* slow high pointer */
    function twoSum($numbers, $target)
    {
        $low = 0;
        $high = count($numbers) - 1;
        while ($low < $high) {
            $sum = $numbers[$low] + $numbers[$high];
            echo $sum, $low, $high, "\n";
            if ($sum == $target) {
                return [$low + 1, $high + 1];
            } else if ($sum > $target) {
                $high--;
            } else {
                $low++;
            }
        }
        return [];
    }
}

