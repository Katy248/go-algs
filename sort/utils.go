package sort

func CopyArray(source []int32) (result []int32) {
	result = make([]int32, len(source))
	copy(result[:], source[:])
	return
}
