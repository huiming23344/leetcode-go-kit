package util

// loop-based comparison 
// typically more efficient as it avoids the overhead of reflection
func IntSliceIsEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for index, val := range(slice1) {
		if val != slice2[index] {
			return false
		}
	}
	return true
}