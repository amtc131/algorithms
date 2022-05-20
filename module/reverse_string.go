package module

// Given a text string, our goal is to write a recursive function (method)
// that is reverses the string characters. For example.
// Example 1
// Input   : ABCDE
// Output  : EDCBA

// Example 2
// Input   : 654A321
// Output  : 123A456
func reversestring(str string, n int) string {
	if n >= 0 {
		return string(str[n]) + reversestring(str, n-1)
	}
	return ""
}

func ReverseStr(str string) string {
	return reversestring(str, len(str)-1)
}
