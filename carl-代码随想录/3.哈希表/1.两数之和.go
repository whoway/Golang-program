func twoSum(nums []int, target int) []int {
	// 注意，这个哈希是存储左边某个元素【最邻近】你当前index的位置
	left_val_new_index_Hash := make(map[int]int)
	var res []int
	for index, val := range nums {
		if 0 != left_val_new_index_Hash[target-val] {
			res = append(res, left_val_new_index_Hash[target-val]-1)
			res = append(res, index)
			return res
		}

		// 巧妙的：加个offerset偏移,这样就不用怕左边某个目标index为0了
		left_val_new_index_Hash[val] = index + 1
	}
	// 给leetcode吃
	return nil
}
