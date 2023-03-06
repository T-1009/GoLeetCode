// https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#%E6%80%9D%E8%B7%AF

package GoLeetCode

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + ((right-left)>>1)
        if nums[mid] < target {
            left = mid + 1
        } else if nums[mid] > target {
            right = mid - 1
        } else {
            return mid
        }
    }
    return -1
}