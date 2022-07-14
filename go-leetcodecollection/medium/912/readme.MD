十大经典排序算法比较
前7基于比较的排序，时间效率极限到O(nlogn)
基础排序

1.冒泡排序-稳定-每轮前部排序-(无序区，有序区)
2.选择排序-不稳定-每轮后部排序-(有序区，无序区)
3.插入排序-稳定-每轮前部倒序插入-(相对有序，无序)
改良排序

4.希尔排序-不稳定
归并、快排、堆排适合数据规模大的排序，渐进性的最佳排序算法
归并并行计算，性能好很多，时间O(n/logn)
5.归并排序-稳定-递归与非递归实现
解决快排的最坏情况-O(n^2)，随机主元
6.快速排序-不稳定
7.堆排序-不稳定
后3基于哈希思想，时间效率到O(n)，空间换时间
空间换时间，适合数据规模较小
有重复数字出现，用计数/基数排序；基数排序是对计数排序的优化，在辅助空间上的优化，基数排序所需辅助空间为：10个该元素空间，计数排序所需辅助空间为：最大数字+1个空间

8.基数排序
无重复数字建议桶排序
9.桶排序
10.计数排序
代码
基于比较的排序算法
基础排序算法
1.冒泡排序


func sortArray(nums []int) []int {
    // 冒泡排序，比较交换，稳定算法，时间O(n^2), 空间O(1)
	// 每一轮遍历，将该轮最大值放到后面，同时小的往前冒
	// 从而形成后部是有序区
	// compare and swap
	for i:=0;i<len(nums);i++ {
		// 适当剪枝，len()-i到最后的部分都是有序区，避免再排
		for j:=1;j<len(nums)-i;j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}
2.选择排序


func sortArray(nums []int) []int {
	// 选择排序，比较交换，不稳定算法，时间O(n^2)，空间O(1)
	// 每一轮遍历，该轮的最小值前挪，从而形成前面部分是有序区
	// compare and swap
	for i:=0;i<len(nums);i++ {
		// 剪枝前面部分，比较后面部分
		for j:=i+1;j<len(nums);j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
3.插入排序


func sortArray(nums []int) []int {
	// 插入排序，比较交换，稳定算法，时间O(n^2)，空间O(1)
	// 0->len方向，每轮从后往前比较，相当于找到合适位置，插入进去
	// 数据规模小的时候，或基本有序，效率高
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > tmp { //左边比右边大
			nums[j+1] = nums[j] //右移1位
			j--                 //扫描前一个数
		}
		nums[j+1] = tmp //添加到小于它的数的右边
	}
	return nums
}
改良排序算法
4.希尔排序


func sortArray(nums []int) []int {
	// 希尔排序，比较交换，不稳定算法，时间O(nlog2n)最坏O(n^2), 空间O(1)
	// 改进插入算法
	// 每一轮按照间隔插入排序，间隔依次减小，最后一次一定是1
	/*
	主要思想：
	设增量序列个数为k，则进行k轮排序。每一轮中，
	按照某个增量将数据分割成较小的若干组，
	每一组内部进行插入排序；各组排序完毕后，
	减小增量，进行下一轮的内部排序。
	*/
	gap := len(nums)/2
	for gap > 0 {
		for i:=gap;i<len(nums);i++ {
			j := i
			for j-gap >= 0 && nums[j-gap] > nums[j] {
				nums[j-gap], nums[j] = nums[j], nums[j-gap]
				j -= gap
			}
		}
		gap /= 2
	}
	return nums
}
5.归并排序


// 递归实现归并算法
func sortArray(nums []int) []int {
	// 归并排序，基于比较，稳定算法，时间O(nlogn)，空间O(logn) | O(n)
	// 基于递归的归并-自上而下的合并，另有非递归法的归并(自下而上的合并)
	// 都需要开辟一个大小为n的数组中转
	// 将数组分为左右两部分，递归左右两块，最后合并，即归并
	// 如在一个合并中，将两块部分的元素，遍历取较小值填入结果集
	// 类似两个有序链表的合并，每次两两合并相邻的两个有序序列，直到整个序列有序
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var l,r,i int
		// 通过遍历完成比较填入res中
		for l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		// 如果left或者right还有剩余元素，添加到结果集的尾部
		copy(res[i:], left[l:])
		copy(res[i+len(left)-l:], right[r:])
		return res
	}
	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}
		// 拆分递归与合并
		// 分割点
		mid := len(nums)/2
		left := sort(nums[:mid])
		right := sort(nums[mid:])
		return merge(left, right)
	}
	return sort(nums)
}

// 非递归实现归并算法
func sortArray(nums []int) []int {
	// 归并排序-非递归实现，利用变量，自下而上的方式合并
	// 时间O(nlogn)，空间O(n)
	if len(nums) <= 1 {return nums}
	merge := func(left, right []int) []int {
		res := make([]int, len(left)+len(right))
		var l,r,i int
		// 通过遍历完成比较填入res中
		for l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				res[i] = left[l]
				l++
			} else {
				res[i] = right[r]
				r++
			}
			i++
		}
		// 如果left或者right还有剩余元素，添加到结果集的尾部
		copy(res[i:], left[l:])
		copy(res[i+len(left)-l:], right[r:])
		return res
	}
	i := 1 //子序列大小初始1
	res := make([]int, 0)
	// i控制每次划分的序列长度
	for i < len(nums) {
		// j根据i值执行具体的合并
		j := 0
		// 按顺序两两合并，j用来定位起始点
		// 随着序列翻倍，每次两两合并的数组大小也翻倍
		for j < len(nums) {
			if j+2*i > len(nums) {
				res = merge(nums[j:j+i], nums[j+i:])
			} else {
				res = merge(nums[j:j+i], nums[j+i:j+2*i])
			}
			// 通过index控制每次将合并的数据填入nums中
			// 重填入的次数和合并及二叉树的高度相关
			index := j
			for _, v := range res {
				nums[index] = v
				index++
			}
			j = j + 2*i
		}
		i *= 2
	}
	return nums
}
6.快速排序


func sortArray(nums []int) []int {
    // 快速排序，基于比较，不稳定算法，时间平均O(nlogn)，最坏O(n^2)，空间O(logn)
	// 分治思想，选主元，依次将剩余元素的小于主元放其左侧，大的放右侧
	// 然后取主元的前半部分和后半部分进行同样处理，直至各子序列剩余一个元素结束，排序完成
	// 注意：
	// 小规模数据(n<100)，由于快排用到递归，性能不如插排
	// 进行排序时，可定义阈值，小规模数据用插排，往后用快排
	// golang的sort包用到了快排
	// (小数，主元，大数)
	var quick func(nums []int, left, right int) []int
	quick = func(nums []int, left, right int) []int {
		// 递归终止条件
		if left > right {
			return nil
		}
		// 左右指针及主元
		i, j, pivot := left, right, nums[left]
		for i < j {
			// 寻找小于主元的右边元素
			for i < j && nums[j] >= pivot {
				j--
			}
			// 寻找大于主元的左边元素
			for i < j && nums[i] <= pivot {
				i++
			}
			// 交换i/j下标元素
			nums[i], nums[j] = nums[j], nums[i]
		}
		// 交换元素
		nums[i], nums[left] = nums[left], nums[i]
		quick(nums, left, i-1)
		quick(nums, i+1, right)
		return nums
	}
	return quick(nums, 0, len(nums)-1)
}
7.堆排序


func sortArray(nums []int) []int {
    // 堆排序-大根堆，升序排序，基于比较交换的不稳定算法，时间O(nlogn)，空间O(1)-迭代建堆
	// 遍历元素时间O(n)，堆化时间O(logn)，开始建堆次数多些，后面次数少 
	// 主要思路：
	// 1.建堆，从非叶子节点开始依次堆化，注意逆序，从下往上堆化
	// 建堆流程：父节点与子节点比较，子节点大则交换父子节点，父节点索引更新为子节点，循环操作
	// 2.尾部遍历操作，弹出元素，再次堆化
	// 弹出元素排序流程：从最后节点开始，交换头尾元素，由于弹出，end--，再次对剩余数组元素建堆，循环操作
	// 建堆函数，堆化
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		// 大顶堆堆化，堆顶值小一直下沉
		for {
			// 左孩子节点索引
			child := root*2 + 1
			// 越界跳出
			if child > end {
				return
			}
			// 比较左右孩子，取大值，否则child不用++
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			// 如果父节点已经大于左右孩子大值，已堆化
			if nums[root] > nums[child] {
				return
			}
			// 孩子节点大值上冒
			nums[root], nums[child] = nums[child], nums[root]
			// 更新父节点到子节点，继续往下比较，不断下沉
			root = child
		}
	}
	end := len(nums)-1
	// 从最后一个非叶子节点开始堆化
	for i:=end/2;i>=0;i-- {
		heapify(nums, i, end)
	}
	// 依次弹出元素，然后再堆化，相当于依次把最大值放入尾部
	for i:=end;i>=0;i-- {
		nums[0], nums[i] = nums[i], nums[0]
		end--
		heapify(nums, 0, end)
	}
	return nums
}
基于哈希思想的排序算法
8.基数排序
先以个位数的⼤小来对数据进⾏排序，接着以十位数的⼤小来对数据进⾏排序，接着以百位数的⼤小......
排到最后，就是一组有序的元素了。不过，他在以某位数进行排序的时候，是⽤用“桶”来排序的。
由于某位数(个位/⼗位....，不是整个数)的⼤小范围为0-9，所以我们需要10个桶，然后把具有相同 数值的数放进同⼀个桶⾥，之后再把桶里的数按照0号桶到9号桶的顺序取出来，这样一趟下来，按照某位数的排序就完成了。
算法效率：
时间复杂度: O(kn)
空间复杂度: O(n+k)
稳定

基数排序转载自：
作者：xilepeng
链接：https://leetcode-cn.com/problems/sort-an-array/solution/golang-by-xilepeng-2/


func radix_sort(li []int) {
	// 一次遍历获取最大值
    max_num := li[0]
    for i := 0; i < len(li); i++ {
        if max_num < li[i] {
            max_num = li[i]
        }
    }
	// 根据最大值的位数确定分几轮基数排序，如234，需要3轮，9仅需要一轮排序
    for j := 0; j < len(strconv.Itoa(max_num)); j++ {
		// 1.每轮排序，先分桶，数据装桶
        bin := make([][]int, 10)
        for k := 0; k < len(li); k++ {
            n := li[k] / int(math.Pow(10, float64(j))) % 10
            bin[n] = append(bin[n], li[k])
        }
		// 2.每轮排序，装完桶后，依次遍历桶，重排数组
        m := 0
        for p := 0; p < len(bin); p++ {
            for q := 0; q < len(bin[p]); q++ {
                li[m] = bin[p][q]
                m++
            }
        }
    }
}

9.桶排序

func sortArray(nums []int) []int {
    // 桶排序，基于哈希思想的外排稳定算法，空间换时间，时间O(n+k)
	// 相当于计数排序的改进版，服从均匀分布，先将数据分到有限数量的桶中，
	// 每个桶分别排序，最后将非空桶的数据拼接起来
	var bucket func(nums []int, bucketSize int) []int
	bucket = func(nums []int, bucketSize int) []int {
		if len(nums) < 2 {
			return nums
		}
		// 获取最大最小值
		minAndMax := func(nums []int) (min, max int) {
			minNum := math.MaxInt32
			maxNum := math.MinInt32
			for i:=0;i<len(nums);i++ {
				if nums[i] < minNum {
					minNum = nums[i]
				}
				if nums[i] > maxNum {
					maxNum = nums[i]
				}
			}
			return minNum, maxNum
		}
		min_, max_ := minAndMax(nums)
		// 定义桶
		// 构建计数桶
		bucketCount := (max_-min_)/bucketSize + 1
		buckets := make([][]int, bucketCount)
		for i:=0;i<bucketCount;i++ {
			buckets[i] = make([]int, 0)
		}
		// 装桶-排序过程
		for i:=0;i<len(nums);i++ {
			// 桶序号
			bucketNum := (nums[i]-min_) / bucketSize
			buckets[bucketNum] = append(buckets[bucketNum], nums[i])
		}
		// 桶中排序
		// 上述装桶完成，出桶填入元素组
		index := 0
		for _, bucket := range buckets {
			sort.Slice(bucket, func(i, j int) bool {
				return bucket[i] < bucket[j]
			})
			for _, num := range bucket {
				nums[index] = num
				index++
			}
		}
		return nums
	}
	// 定义桶中的数量
	var bucketSize int = 2
	return bucket(nums, bucketSize)
}
10.计数排序算法


func sortArray(nums []int) []int {
    // 计数排序，基于哈希思想的稳定外排序算法，空间换时间，时间O(n)，空间O(n)
	// 数据量大时，空间占用大
	// 空间换时间，通过开辟额外数据空间存储索引号记录数组的值和数组额个数
	// 思路：
	// 1.找出待排序的数组的最大值和最小值
	// 2.创建数组存放各元素的出现次数，先于[min, max]之间
	// 3.统计数组值的个数
	// 4.反向填充数组，填充时注意,num[i]=j+min，
	// j-前面需要略过的数的个数，两个维度，依次递增的数j++，一个是重复的数的计数j-不变
	if len(nums) == 0 {
		return nums
	}
	// 获取最大最小值
	minAndMax := func(nums []int) (min,max int) {
		minNum := math.MaxInt32
		maxNum := math.MinInt32
		for i:=0;i<len(nums);i++ {
			if nums[i] < minNum {
				minNum = nums[i]
			}
			if nums[i] > maxNum {
				maxNum = nums[i]
			}
		}
		return minNum, maxNum
	}
	min_, max_ := minAndMax(nums)
	// 中转数组存放遍历元素
	// 空间只需要min-max
	tmpNums := make([]int, max_-min_+1)
	// 遍历原数组
    for i:=0;i<len(nums);i++ {
		tmpNums[nums[i]-min_]++
	}
	// 遍历中转数组填入原数组
	j := 0
	for i:=0;i<len(nums);i++ {
		// 如果对应数字cnt=0，说明可以计入下一位数字
		for tmpNums[j] == 0 {
			j++
		}
		// 填入数字
		nums[i] = j + min_
		// 填一个数字，对应数字cnt--
		tmpNums[j]--
	}
	return nums
}