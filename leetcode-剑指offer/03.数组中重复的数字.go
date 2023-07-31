func findRepeatNumber(nums []int) int {
	myhash := make(map[int]int)
	for _, val := range nums {
		myhash[val]++
		// fmt.Println(myhash[val]);
	}

	for _, val := range nums {
		if myhash[val] > 1 {
			return val
		}
	}
	return -1
}
