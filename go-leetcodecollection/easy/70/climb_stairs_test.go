package _70

import "testing"

func TestClimbStairs(t *testing.T) {
	n := 3
	target := 3
	res := climbStairs(n)
	if res != target {
		t.Fatalf("res=%d 希望 %d\n", res, target)
	}
}
