package main

import (
	"fmt"
)

func totalFruit(fruits []int) int {
    slowIndex, fastIndex, res := 0, 0, -1
	ln, rn := fruits[slowIndex], fruits[fastIndex]
    for fastIndex < len(fruits) {
		if fruits[fastIndex] == ln || fruits[fastIndex] == rn {
			res = max(res, fastIndex - slowIndex + 1)
			fastIndex ++
	    } else {	//如果遇到第三种，把慢指针移动到快指针前一步，该步的水果种类必然不同于快指针，此时慢指针慢慢回退齐所有的连续同类。
			slowIndex = fastIndex - 1
			ln = fruits[slowIndex]
			for slowIndex >= 1 && fruits[slowIndex-1] == ln {
				slowIndex -- // 回溯到第一个同类 
			}
			rn = fruits[fastIndex]
			res = max(res, fastIndex - slowIndex + 1)
		}
    }
    return res
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
	a :=  []int{3,3,3,1,2,1,1,2,3,3,4}

	res := totalFruit(a)
	
	fmt.Println(res)

}