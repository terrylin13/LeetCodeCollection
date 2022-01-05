package _797

func allPathsSourceTarget(graph [][]int) [][]int {
	source, dest := 0, len(graph)-1
	ans := [][]int{}
	paths := []int{}

	var dfs func(vert int)

	dfs = func(vert int) {
		paths = append(paths, vert)

		if vert == dest {
			temp := make([]int, len(paths))
			copy(temp, paths)
			ans = append(ans, temp)
		} else {
			for _, i := range graph[vert] {
				dfs(i)
			}
		}

		paths = paths[:len(paths)-1]
	}

	dfs(source)
	return ans
}
