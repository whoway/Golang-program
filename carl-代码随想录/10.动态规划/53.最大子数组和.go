func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	mylen := len(nums)
	ret := nums[0]
	dp[0] = nums[0]
	for i := 1; i < mylen; i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}

		if ret < dp[i] {
			ret = dp[i]
		}
	}
	return ret
}
