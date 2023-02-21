package GoLeetCode

func preimageSizeFZF(k int) int {
	base := 0
	for base < k {
		base = 5*base + 1
	}
	for k > 0 {
		base = (base - 1) / 5
		if k/base == 5 {
			return 0
		}
		k %= base
	}
	return 5
}

// a = [458,32,72,64,5,1615,1466,997,332,310,46,51,77,95,797,55,134,135,605,621,62,70,120,124,198,23,53,148,327,654,21,60,143,233,342]