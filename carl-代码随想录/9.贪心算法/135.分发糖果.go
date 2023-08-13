
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func candy(ratings []int) int {
	arr := make([]int, len(ratings))
	arr[0] = 1
	mylen := len(ratings)

	// 先满足规则1: 右边的比相邻的左边的分数高, 最优的方法
	for i := 1; i < mylen; i++ {
		if ratings[i] > ratings[i-1] {
			arr[i] = arr[i-1] + 1
		} else {
			//也就是默认的至少每个人得有个糖果
			arr[i] = 1
		}
	}

	// 先满足规则2: 左边的比相邻的右边的分数高, 最优的方法是取,满足规则2和规则1的最大值!
	//[因为必须都得满足]
	for i := mylen - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			arr[i] = max(arr[i+1]+1, arr[i])
		}
	}
	result := 0
	for i := 0; i < mylen; i++ {
		result += arr[i]
	}
	return result
}
