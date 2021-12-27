package _322

import "testing"

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	// amount := 100
	amount := 37
	t.Log(coinChange2(coins, amount))
}
