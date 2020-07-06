<?php
class Solution
{

    /**
     * @param Integer[][] $obstacleGrid
     * @return Integer
     */
    function uniquePathsWithObstacles($obstacleGrid)
    {
        $lengthX = count($obstacleGrid);
        $lengthY = count($obstacleGrid[0]);
        if ($obstacleGrid[0][0] == 1 || $obstacleGrid[$lengthX - 1][$lengthY - 1] == 1) {
            return 0;
        }

        //initial
        $dp = [];
        $dp[0][0] = 1;
        for ($x = 1; $x < $lengthX; ++$x) {
            $dp[$x][0] = ($obstacleGrid[$x][0] == 1 || $dp[$x - 1][0] == 0) ? 0 : 1;
        }

        for ($y = 1; $y < $lengthY; ++$y) {
            $dp[0][$y] = ($obstacleGrid[0][$y] == 1 || $dp[0][$y - 1] == 0) ? 0 : 1;
        }

        for ($x = 1; $x < $lengthX; $x++) {
            for ($y = 1; $y < $lengthY; $y++) {
                if ($obstacleGrid[$x][$y] == 0) {
                    $dp[$x][$y] = $dp[$x - 1][$y] + $dp[$x][$y - 1];
                } else {
                    $dp[$x][$y] = 0;
                }
            }
        }
        return $dp[$lengthX - 1][$lengthY - 1];
    }
}
