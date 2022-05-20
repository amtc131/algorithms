package module

func ReverseArray(arr []int) {
	reverseAray(arr, 0, len(arr)-1)
}

func reverseAray(arr []int, i int, l int) {
	if i < l {
		var temp = arr[i]
		arr[i] = arr[l]
		arr[l] = temp
		reverseAray(arr, i+1, l-1)
	}
}
