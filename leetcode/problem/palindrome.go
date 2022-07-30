package problem

import "strconv"

func reverse(s string) string {
	ss := []rune(s)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return string(ss)
}
func isPalindrome(x int) bool {
	y := strconv.Itoa(x)
	return y == reverse(y)
}
