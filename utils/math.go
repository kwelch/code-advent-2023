package utils

func SumValues(values []int) int {
	var total int = 0
	for _, value := range values {
		total += value
	}
	return total
}

func PowerValues(values []int) int {
	var total int = 0
	for i, value := range values {
		if i == 0 {
			total = value
			continue
		}
		total *= value
	}
	return total
}

func LowestValue(values []int) int {
	var lowest = 0
	for i, value := range values {
		if i == 0 {
			lowest = value
			continue
		}
		if value < lowest {
			lowest = value
		}
	}
	return lowest
}
