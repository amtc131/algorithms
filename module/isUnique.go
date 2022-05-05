package module

// Is Unique implement an algoritm to determine if a string has all unique characters.
// What if you cannot additional data structure
func IsUnique(str string) bool {
	c := make([]bool, 128)

	for i := 0; i < len(str); i++ {
		val := str[i]
		if c[val] {
			return false
		}
		c[val] = true
	}

	return true
}

//	k := 1
//	for i := 0; i < len(str); i++ {
//		for j := 1; j < len(str); j++ {
//			if string(str[i]) == string(str[j]) {
//				k++
//			}
//		}
//		if k <= 1 {
//			k = 0
//		} else {
//			return false
//		}
//	}
//	return true
