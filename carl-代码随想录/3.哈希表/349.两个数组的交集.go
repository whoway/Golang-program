func intersection(nums1 []int, nums2 []int) []int {
	myhash := make(map[int]int)
	// 用于标记，防止出现重复放入
	tag := make(map[int]bool)
	for _, val := range nums1 {
		myhash[val]++
		tag[val] = true
	}

	var res []int
	for _, val := range nums2 {
		if myhash[val] > 0 && true == tag[val] {
			res = append(res, val)
			tag[val] = false
		}
	}

	return res
}
