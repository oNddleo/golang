package problem

func longestPalindrome(s string) string {
	longest := ""
	for i := 0; i < len(s); i++ {
		// check longest is even string or odd string
		odd := expand(s, i, i)
		even := expand(s, i, i+1)
		longer := ""
		if len(odd) > len(even) {
			longer = odd
		} else {
			longer = even
		}

		if len(longer) > len(longest) {
			longest = longer
		}
	}
	return longest
}

func expand(s string, start int, end int) string {
	for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
		start -= 1
		end += 1
	}
	return s[start+1 : end]
}
