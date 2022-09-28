package sort

func BubbleSort(num []int) []int {
	for i := 0; i < len(num); i++ {
		swapHappened := false
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				swapHappened = true
				num[i], num[j] = num[j], num[i]
			}
		}
		if !swapHappened {
			return num
		}
	}
	return num
}
