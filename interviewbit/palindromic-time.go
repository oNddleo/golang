package interviewbit

import (
	"fmt"
	"strconv"
	"strings"
)

func Palindromic(A string) int {
	s := strings.Split(A, ":")
	hour, _ := strconv.Atoi(s[0])
	minute, _ := strconv.Atoi(s[1])
	fmt.Print(hour, minute, "\n")
	ans := 0
	for {
		if hour/10 == minute%10 && minute/10 == hour%10 {
			break
		}
		minute++
		ans++
		if minute%60 == 0 {
			hour++
			minute = minute % 60
		}
		if hour%24 == 0 {
			hour = hour % 24
		}
		fmt.Println(hour, minute)
	}
	return ans
}
