package main

import (
	s "algs/sort"
	"fmt"
	"math/rand"
	"time"
)

var (
	ArrayLength          = 100
	ArrayElementMaxValue = 100
)

func main() {
	array := generateArray()
	fmt.Println("Source array:\n", array)

	printSortedArray("Exchange sort", &s.ExchangeSorter{}, array)

	printSortedArray("Selection sort", &s.SelectionSorter{}, array)

	printSortedArray("Bubble sort", &s.BubbleSorter{}, array)

	printSortedArray("Insertion sort", &s.InsertionSorter{}, array)

	printSortedArray("Shaker sort", &s.ShakerSorter{}, array)

	printSortedArray("Comb sort", &s.CombSorter{}, array)

	printSortedArray("Shell sort", s.NewShellSorter(), array)
}

func printSortedArray(title string, sorter s.Sorter, array []int32) {

	println()
	println(title, ":")
	start := time.Now()
	sortedArray := sorter.Sort(array)
	takenTime := time.Since(start)
	fmt.Println("Sorted array (with time", takenTime, ")\n", sortedArray)
}

func generateArray() []int32 {
	result := make([]int32, ArrayLength)
	for i := range result {
		result[i] = rand.Int31n(int32(ArrayElementMaxValue))
	}
	return result
}
