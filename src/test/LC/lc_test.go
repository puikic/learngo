package LC

import "testing"

func maxSubArray(nums []int) (int, []int) {
	res := nums[0]
	dp := nums[0]
	currentStart := 0
	maxStart, maxEnd := 0, 0

	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp += nums[i]
		} else {
			dp = nums[i]
			currentStart = i
		}

		if dp > res {
			res = dp
			maxStart = currentStart
			maxEnd = i
		}
	}

	return res, nums[maxStart : maxEnd+1]
}

func fmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func TestLC53(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	k, s := maxSubArray(nums)
	t.Log(k, s)
}
