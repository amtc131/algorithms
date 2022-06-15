package binarySearchTree

//Go implementation of Iterative Binary search
func BinarySearch(arr []int, target int) int {
	// Returns index of x if it is present in arr[],
	// else return -1
	l, r := 0, len(arr)-1

	for l <= r {
		m := l + (r-l)/2

		// Check if x is present at mid
		if arr[m] == target {
			return m
		}

		// If x greater, ignore left half
		if arr[m] < target {
			l = m + 1
			// If x is smaller, ignore right half
		} else {
			r = m - 1
		}
	}

	// if we reach here, then element was
	// not present
	return -1
}
