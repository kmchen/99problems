/*
2 solutions available:

1. Sort the string
2. Use a map to count the occurence of each unique character
*/
package main

import (
	"fmt"
)

func anagram(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	charArray1 := make(map[uint8]int)
	charArray2 := make(map[uint8]int)

	for i, _ := range str1 {
		charArray1[str1[i]]++
		charArray2[str2[i]]++
	}

	for key, value := range charArray1 {
		if value != charArray2[key] {
			return false
		}
	}
	return true
}

func main() {
	//char str1[] = "geeksforgeeks"
	//char str2[] = "forgeeksgeeks";
	fmt.Println(anagram("geeksforgeeks", "forgeeksgeeks"))
}
