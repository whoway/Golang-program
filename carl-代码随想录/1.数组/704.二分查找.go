func search(nums []int, target int) int {
	// 元素不重复的二分查找
	if 0 == len(nums) {
		return -1
	}

	// 我采用的【左闭右开）的写法
	left := 0
	right := len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return -1
}
