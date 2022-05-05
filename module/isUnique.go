package module

func IsUnique(str string) bool {

	k := 1

	for i := 0; i < len(str); i++ {
		for j := 1; j < len(str); j++ {
			if string(str[i]) == string(str[j]) {
				k++
			}
		}
		if k <= 1 {
			k = 0
		} else {
			return false
		}
	}
	return true
}
