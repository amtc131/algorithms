package module

func ReverseNumber(number int) int {

	if number < 0 {
		return -reverse(-number, 0)
	}

	return reverse(number, 0)
}

func reverse(n int, r int) int {
	if n > 0 {
		//fmt.Println(" ", n, " - R: ", r)
		return reverse(n/10, r*10+(n%10))
	}
	return r
}
