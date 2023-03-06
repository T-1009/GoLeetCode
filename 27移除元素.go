package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
    slowIndex := 0
    for fastIndex, value := range nums {
        if value != val {
            nums[slowIndex] = nums[fastIndex]
            slowIndex ++
        }
    }
    return slowIndex
}

func main() {
	nums := [...]int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	res := removeElement(nums[:], val)
	fmt.Println(res, nums)
}