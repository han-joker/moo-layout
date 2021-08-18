package toolm

func IntSliceContains(value int, container []int) bool {
	for _, v := range container {
		if v == value {
			return true
		}
	}
	return false
}
