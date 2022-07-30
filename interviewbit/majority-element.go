package interviewbit

import "sort"

func MajorityElement(A []int) int {
	sort.Ints(A)
	return A[len(A)/2]
}
