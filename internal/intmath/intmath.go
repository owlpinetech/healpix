package intmath

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Clamp(val int, low int, high int) int {
	if val < low {
		return low
	} else if val > high {
		return high
	} else {
		return val
	}
}

func Wrap(val int, low int, high int) int {
	if val < low {
		return val + (high + 1 - low)
	} else if val > high {
		return val - (high + 1)
	} else {
		return val
	}
}

// Pow returns x**y, the base-x exponential of y. Negative exponents return 0.
func Pow(base int, exp int) int {
	if exp < 0 {
		return 0
	}

	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// Exp2 returns 2**x, the base-2 exponential of x. Negative exponents return 0.
func Exp2(exp int) int {
	if exp < 0 {
		return 0
	}

	return 1 << exp
}

// Log2 returns x such that 2**x <= num and 2**(x+1) > num. Negative numbers panic.
func Log2(num int) int {
	if num < 0 {
		panic("intmath: Log2 argument must be 0 or greater")
	}
	// TODO: using more bit twiddling or different comparisons this can probably be made faster
	res := 0
	for num >>= 1; num != 0; num >>= 1 {
		res += 1
	}
	return res
}
