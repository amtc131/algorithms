package module

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(number []int) int {
	s := 0
	for _, i := range number {
		s += i
	}
	return s
}

// recursive
// if len(number) == 0 {
//	return 0
//}
//return number[0] + Sum(number[1:])*/
