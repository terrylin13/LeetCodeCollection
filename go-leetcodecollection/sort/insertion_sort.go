package sort

// 插入排序，a 表示数组，n 表示数组大小
func insertionSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1

		// 查找插入的位置
		for ; j >= 0; j-- {
			if a[j] > value {
				// 数据移动
				a[j+1] = a[j]
			} else {
				break
			}
		}

		// 插入数据
		a[j+1] = value
	}
}
