package problem

func isValid(s string) bool {
	pair := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		n := len(stack)
		if n > 0 && pair[stack[n-1]] == s[i] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0

}
