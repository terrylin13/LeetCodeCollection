package _188

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	target := 7
	res := maxProfit2(k, prices)
	str := fmt.Sprintf("target:%d,res:%d", target, res)
	t.Log(str)
}
