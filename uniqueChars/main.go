package main

import (
	"fmt"
)

func uniqueChars(str string) bool {
	result := 0
	for _, c := range str {
		if (result & (1 << uint(c-'a'))) > 0 {
			return false
		}
		result |= 1 << uint(c-'a')
	}
	return true
}

func main() {
	str := "abcdefghijklmnopqrst"
	fmt.Println("string : ", str, uniqueChars(str))
}
