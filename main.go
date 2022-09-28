package main

import (
	"fmt"
	"github.com/PavelDonchenko/algorithms/sort"
)

func main() {
	num := []int{2, 3, 45, 5, 6, 7, 8, 9, 9, 0, 0, 8, 6, 5}

	fmt.Println(sort.BubbleSort(num))
	fmt.Println(sort.MergeSort(num))
	fmt.Println(sort.QuickSort(num))
}
