package module

import (
	"strings"
)

func Urlify(str string) string {
	s := strings.Count(str, " ")
	l := len(str) + (s * 2)
	n := make([]byte, l)
	i, j := 0, 0
	//j := 0
	for i < len(str) {
		if str[i] == ' ' {
			n[j] = '%'
			n[j+1] = '2'
			n[j+2] = '0'
			j += 3
		} else {
			n[j] = str[i]
			j++
		}
		i++
	}
	return string(n)
}
