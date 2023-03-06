package GoLeetCode

func abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}

func sortedSquares(nums []int) []int {
    left, right := 0, len(nums)-1
    res := []int{}
    for left <= right {
        if abs(nums[left]) < abs(nums[right]) {
            res = append([]int{nums[right] * nums[right]}, res...)
            right --
        } else {
            res = append([]int{nums[left] * nums[left]}, res...)
            left ++
        }
    }
    return res
}