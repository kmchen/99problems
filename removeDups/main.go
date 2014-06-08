package main

import (
	"fmt"
)

func removeDups(array []byte) []byte {

	if len(array) < 2 {
		return []byte(``)
	}
	if array == nil {
		return []byte(``)
	}
	for i := 1; i < len(array); i++ {
		for j := 0; j <= i-1; j++ {
			if array[i] == array[j] {
				array[i] = 0
			}
		}
	}

	return array
}

func main() {
	fmt.Println(string(removeDups([]byte("aaabb"))))
}
