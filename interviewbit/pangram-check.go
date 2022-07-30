package interviewbit

import "unicode"

func Pangram(A []string) int {
	check := make(map[rune]bool)
	for i := 0; i < len(A); i++ {
		for _, char := range A[i] {
			check[unicode.ToLower(char)] = true
		}
	}
	count := 0
	for _, v := range check {
		if v {
			count++
		}
	}
	if count == 26 {
		return 1
	}
	return 0
}
