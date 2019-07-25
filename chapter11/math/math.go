package math

// comments immediately above a function will print with go doc
// Can have multiple lines without any empty lines

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
	if len(xs) < 1 {
		return -1	// can I return an error type??
	}

	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

// Finds the minimum of a series of numbers
func Min(xs []float64) float64 {
	if len(xs) < 1 {
		return -1	// can I return an error type??
	}
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// Finds the maximum of a series of numbers
func Max(xs []float64) float64 {
	if len(xs) < 1 {
		return -1	// can I return an error type??
	}
	max := xs[0]
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

