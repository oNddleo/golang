/*
	Given two integers A and B, where A is the first element of the sequence then find Bth element of the sequence.
	If the kth element of the sequence is X then k+1th element calculated as:
	if X is even then next element is X/2.
	else next element is 3×X + 1.
*/
package interviewbit

import "fmt"

func Solve(A int, B int) int64 {
	result := A
	for i := 1; i <= B-1; i++ {
		if result%2 == 0 {
			result = result / 2
		} else {
			result = (3 * result) + 1
		}
		fmt.Print(result, " -> ")
	}
	return int64(result)
}
