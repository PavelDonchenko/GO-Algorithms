package sort

import (
	"math/rand"
	"testing"
)

func BenchmarkBubbleSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(rand.Perm(100))
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(rand.Perm(10000))
	}
}

func BenchmarkMergeSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(rand.Perm(100))
	}
}
func BenchmarkMergeSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(rand.Perm(10000))
	}
}

func BenchmarkQuickSort100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(rand.Perm(100))
	}
}
func BenchmarkQuickSort10000(b *testing.B) {
	for n := 0; n < b.N; n++ {
		QuickSort(rand.Perm(10000))
	}
}
