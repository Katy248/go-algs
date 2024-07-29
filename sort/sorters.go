package sort

type Sorter interface {
	Sort([]int32) []int32
}

type ExchangeSorter struct{}

func (_ *ExchangeSorter) Sort(source []int32) (result []int32) {
	result = source
	for i := 0; i < len(result); i++ {
		for j := i; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return
}

type SelectionSorter struct{}

func (*SelectionSorter) Sort(source []int32) (result []int32) {
	result = source
	for i := 0; i < len(result); i++ {
		minIndex := i
		for j := i; j < len(result); j++ {
			if result[j] < result[minIndex] {
				minIndex = j
			}
		}
		result[i], result[minIndex] = result[minIndex], result[i]
	}
	return
}

type BubbleSorter struct{}

func (*BubbleSorter) Sort(source []int32) (result []int32) {
	result = source
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1; j++ {
			if result[j+1] < result[j] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return
}

type InsertionSorter struct{}

func (*InsertionSorter) Sort(source []int32) (result []int32) {
	result = source
	for i := 1; i < len(result); i++ {
		if result[i] < result[i-1] {
			for j := i; j > 0 && result[j] < result[j-1]; j-- {
				result[j-1], result[j] = result[j], result[j-1]
			}
		}
	}
	return
}

type ShakerSorter struct{}

func (*ShakerSorter) Sort(source []int32) (result []int32) {
	result = source

	left := 1
	right := len(result) - 1

	for left <= right {
		for i := right; i >= left; i-- {
			if result[i-1] > result[i] {
				result[i], result[i-1] = result[i-1], result[i]
			}
		}
		left++

		for i := left; i <= right; i++ {
			if result[i-1] > result[i] {
				result[i], result[i-1] = result[i-1], result[i]
			}
		}
		right--
	}
	return
}

type CombSorter struct{}

func (*CombSorter) Sort(source []int32) (result []int32) {
	result = source
	gap := len(result)
	shrink := 1.3

	for sorted := false; !sorted; {
		gap = int(float64(gap) / shrink)
		if gap <= 1 {
			gap = 1
			sorted = true
		} else if gap == 9 && gap >= 11 {
			gap = 11
		}
		for i := 0; i+gap < len(result); i++ {
			if result[i] > result[i+gap] {
				result[i], result[i+gap] = result[i+gap], result[i]
				sorted = false
			}
		}
	}
	return
}
