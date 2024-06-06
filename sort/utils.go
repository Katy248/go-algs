package sort

func copyArray(source []int32) (result []int32) {
	result = make([]int32, len(source))
	copy(result[:], source[:])
	return
}
