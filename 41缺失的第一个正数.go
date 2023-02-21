package GoLeetCode

func firstMissingPositive(nums []int) int {
    for i:=0; i<len(nums); i++ {
        if nums[i] <= 0 {
            nums[i] = len(nums) + 1
        }
    }
    for i:=0; i<len(nums); i++ {
        num := abs(nums[i])
        if num <= len(nums) {
            nums[num-1] = -abs(nums[num-1])
        }
    }
    for i:=0; i<len(nums); i++ {
        if nums[i] > 0 {
            return i+1
        }
    }
    return len(nums) + 1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}