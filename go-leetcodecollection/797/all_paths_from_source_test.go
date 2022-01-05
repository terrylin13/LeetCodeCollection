package _797

import "testing"

func TestAllPathsSourceTarget(t *testing.T) {
	//[[1,2],[3],[3],[]]
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	t.Log(allPathsSourceTarget(graph))
}
