package module

// In this problem, you have to implement the void
// maxMin(int[] arr) method. This will re-arrange the
// elements of a sorted array in such a way that the
// first position will have the largest number,
// the second will have the smallest, the third will have
// the second-largest, and so on.
func MaxMin(arr []int) {

	min := 0
	max := len(arr) - 1
	flag := true
	tmp := make([]int, len(arr))

	_ = copy(tmp, arr)

	for i := 0; i < len(arr); i++ {
		if flag {
			arr[i] = tmp[max]
			max--
		} else {
			arr[i] = tmp[min]
			min++
		}
		flag = !flag
	}
}
