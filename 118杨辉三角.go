package GoLeetCode

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}
	res := [][]int{}
	for j := 1; j <= numRows; j++ {
		stack := make([]int, j)
		stack[0] = 1
		for i := 1; i < j; i++ {
			stack[i] = stack[i-1] * (j - i) / i
		}
		res = append(res, stack)
	}
	return res
}
