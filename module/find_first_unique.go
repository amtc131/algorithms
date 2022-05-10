package module

// In this problem, you have to implement the int findFirstUnique(int[] arr)
// method that will look for a first unique integer,
// which appears only once in the whole array.
// The function returns -1 if no unique number is found.
func FindFirstUnique(arr []int) int {
	isRepeated := false

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] && i != j {
				isRepeated = true
			}
		}
		if isRepeated == false {
			return arr[i]
		} else {
			isRepeated = false
		}
	}

	return -1
}
