package GoLeetCode

func calculate(s string) int {
    i, stack, sign, res := 0, []int{}, 1, 0
    for i < len(s) {
        if s[i] == ' ' {
            i ++
        } else if s[i] >= '0' && s[i] <= '9' {
            base, v := 10, int(s[i]-'0')
            for i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
                v =  v*base + int(s[i+1] - '0')
                i ++
            }
            res += v * sign
            i ++
        } else if s[i] == '-' {
            sign = -1
            i ++
        } else if s[i] == '+' {
            sign = 1
            i ++
        } else if s[i] == '(' {
            stack = append(stack, res)
            stack = append(stack, sign)
            res = 0
            sign = 1
            i ++
        } else if s[i] == ')' {
            /* 模拟栈 */
            signn := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            ress := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            res = res * signn + ress
            i ++
        }
    }
    return res
}