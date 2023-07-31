func sortedSquares(nums []int) []int {
	// 指定初始化的slice的长度的方法
	res := make([]int, len(nums))
	for index, val := range nums {
		res[index] = val * val
	}

	// 使用sort包来排序-我们的元素是int，所以采用Int，然后数目多，所以是s
	sort.Ints(res)
	return res
}
