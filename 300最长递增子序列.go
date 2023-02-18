package GoLeetCode

// O(n log n)

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, v := range nums {
		t := searchFirstGreaterElement(dp, v)
		if t == -1 {
			dp = append(dp, v)
		} else {
			dp[t] = v
		}
	}
	return len(dp)
}

func searchFirstGreaterElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)
		if nums[mid] >= target {
			if (mid == 0) || (nums[mid-1] < target) { // 找到第一个大于等于 target 的元素
				return mid
			}
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
