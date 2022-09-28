package sort

func QuickSort(num []int) []int {
	if len(num) < 1 {
		return num
	}
	pivot := num[len(num)-1]
	var left []int
	var right []int

	for i := 0; i < len(num)-1; i++ {
		if num[i] < pivot {
			left = append(left, num[i])
		} else {
			right = append(right, num[i])
		}
	}
	result := append(left, pivot)
	result = append(result, right...)
	return result
}
