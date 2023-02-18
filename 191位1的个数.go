package GoLeetCode

/*
num & (num-1) 可以去除num的最低位1
*/

func hammingWeight(num uint32) int {
    res := 0
    for num != 0 {
        num &= (num-1)
        res ++
    }
    return res
}
