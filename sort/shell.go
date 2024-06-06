package sort

type ShellSorter struct {
	Step func(int) int
}

func NewShellSorter() *ShellSorter {
	return &ShellSorter{
		Step: ShellStep,
	}
}

func (s *ShellSorter) Sort(source []int32) (result []int32) {
	result = copyArray(source)

	for d := s.Step(len(result)); d != 0; d = s.Step(d) {
		for i := d; i < len(result); i += d {
			if result[i] < result[i-d] {
				result[i-d], result[i] = result[i], result[i-d]
			}
		}
	}
	// return
	sorter := InsertionSorter{}
	return sorter.Sort(result)
}

func ShellStep(step int) int {
	return step / 2
}
