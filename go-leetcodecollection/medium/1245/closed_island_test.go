package _1245

import "testing"

func TestClosedIsland(t *testing.T) {
	grid := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0},
	}
	target := 2
	res := closedIsland(grid)
	t.Logf("target: %d,res: %d", target, res)
}
