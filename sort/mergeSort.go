package sort

func MergeSort(num []int) []int {
	if len(num) <= 1 {
		return num
	}
	middle := len(num) / 2

	return merge(MergeSort(num[middle:]), MergeSort(num[:middle]))

}

func merge(a, b []int) []int {
	var sorted []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		sorted = append(sorted, a[i])
	}
	for ; j < len(b); j++ {
		sorted = append(sorted, b[j])
	}
	return sorted
}
