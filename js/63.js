/**
 * @param {number[][]} obstacleGrid
 * @return {number}
 */
var uniquePathsWithObstacles = function (obstacleGrid) {
    const lengthX = obstacleGrid.length;
    const lengthY = obstacleGrid[0].length;
    if (obstacleGrid[0][0] == 1) {
        return 0;
    }

    //initial
    let dp = new Array(lengthX);
    for (let i = 0; i < lengthX; i++) {
        dp[i] = new Array(lengthY).fill(0);
    }
    dp[0][0] = 1;
    for (let x = 1; x < lengthX; x++) {
        if (obstacleGrid[x][0] != 1) {
            dp[x][0] = dp[x - 1][0];
        }
    }
    for (let y = 1; y < lengthY; y++) {
        if (obstacleGrid[0][y] != 1) {
            dp[0][y] = dp[0][y - 1];
        }
    }

    //dp
    for (let x = 1; x < lengthX; x++) {
        for (let y = 1; y < lengthY; y++) {
            if (obstacleGrid[x][y] != 1) {
                dp[x][y] = dp[x - 1][y] + dp[x][y - 1];
            }
        }
    }

    return dp[lengthX - 1][lengthY - 1];
};
