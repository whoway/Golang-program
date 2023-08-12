func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLIS(nums []int) int {
	mylen := len(nums)
	if 1 == mylen {
		return 1
	}
	dp := make([]int, len(nums))
	for i := 0; i < mylen; i++ {
		dp[i] = 1
	}

	result := 1
	for index := 1; index < mylen; index++ {
		for prev := index - 1; prev >= 0; prev-- {
			if nums[prev] < nums[index] {
				dp[index] = max(dp[index], dp[prev]+1)
				result = max(result, dp[index])
			}
		}
	}
	return result
}
