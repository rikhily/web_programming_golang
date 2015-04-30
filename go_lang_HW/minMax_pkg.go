package min_max

// Finds the minimum value 
func Min(val []float64) float64 {
	min := val[0]
	for _, value := range val {
		if value < min {
			min = value
		}
	}
	return min
}

// Finds the maximum value 
func Max(val []float64) float64 {
	max := val[0]
	for _, value := range val {
		if value > max {
			max = value
		}
	}
	return max
}
