package module

type flt func(int) bool

func Filter(sl []int, f flt) (yes, no []int) {
	for _, val := range sl {
		if f(val) {
			yes = append(yes, val)
		} else {
			no = append(no, val)
		}
	}
	return
}

func IsEvent(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
