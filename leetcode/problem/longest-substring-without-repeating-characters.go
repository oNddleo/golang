package problem

func lengthOfLongestSubstrings(s string) int {
	m, max, left := make(map[rune]int), 0, 0
	for index, c := range s {

		if _, ok := m[c]; ok == true && m[c] >= left {
			if index-left > max {
				max = index - left
			}
			left = m[c] + 1
		}
		m[c] = index
	}

	if len(s)-left > max {
		max = len(s) - left
	}
	return max
}
