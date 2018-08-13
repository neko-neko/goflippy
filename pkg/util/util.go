package util

// Min returns minimum value
func Min(x, y int64) int64 {
	if x < y {
		return x
	}

	return y
}

// Max returns maximum value
func Max(x, y int64) int64 {
	if x < y {
		return y
	}

	return x
}
