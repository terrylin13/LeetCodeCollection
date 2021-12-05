package _63

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	m := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	t.Log(UniquePathsWithObstacles(m))
}
