package main

import (
	conf "algs/cmd"
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
	array := generateArray(conf.ParseConfig())
	fmt.Println("Source array:\n", array)

	sorters := map[string]s.Sorter{
		"Exchange sort":  &s.ExchangeSorter{},
		"Selection sort": &s.SelectionSorter{},
		"Bubble sort":    &s.BubbleSorter{},
		"Insertion sort": &s.InsertionSorter{},
		"Shaker sort":    &s.ShakerSorter{},
		"Combo sort":     &s.CombSorter{},
		"Shell sort":     s.NewShellSorter(),
	}

	for name, sorter := range sorters {
		printSortedArray(name, sorter, array)
	}
}

func printSortedArray(title string, sorter s.Sorter, array []int32) {

	fmt.Println()
	fmt.Printf("%s:\n", title)
	arrCopy := s.CopyArray(array)
	start := time.Now()
	sortedArray := sorter.Sort(arrCopy)
	takenTime := time.Since(start)
	fmt.Println("Sorted array (with time", takenTime, ")\n", sortedArray)
}

func generateArray(config conf.Config) []int32 {
	result := make([]int32, config.ArrayLength)
	for i := range result {
		result[i] = rand.Int31n(int32(config.ArrayMaxElementSize))
	}
	return result
}
