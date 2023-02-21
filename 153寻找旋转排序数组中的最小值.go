package GoLeetCode

func findMin(nums []int) int {
    left, right, mid := 0, len(nums)-1, 0
    for left < right {
        if nums[left] < nums[right] {
            return nums[left]
        }
        mid = left + ((right - left) >> 1)
        if nums[mid] >= nums[left] {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return nums[left]
}