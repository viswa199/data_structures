package pattern

import (
	"fmt"
)

func hash(s string) uint8 {
	var res uint8
	for i := 0; i < len(s); i++ {
		res = res + s[i]
	}
	return res
}

func RabinKarp(text, pattern string) []int {
	var res []int
	for i := 0; i <= len(text)-len(pattern); i++ {
		if hash(pattern) == hash(text[i:i+len(pattern)]) {
			k := i
			for j := 0; j < len(pattern); j++ {
				if pattern[j] != text[k] {
					fmt.Println(j, k)
					break
				}
				k += 1
			}
			res = append(res, i)
		}
	}
	return res
}