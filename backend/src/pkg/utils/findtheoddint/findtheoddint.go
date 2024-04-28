package findtheoddint

func FindOddInt(inputSlice []int) int {
	result := 0

	for _, number := range inputSlice {
		result ^= number
	}

	return result
}
