package main

import "fmt"

func main() {
	input := struct {
		s string
		k int
	}{"Hello how are you Contestant", 4}
	output := "Hello how are you"
	if truncateSentence(input.s, input.k) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	truncate sentence when kth ' ' present
	time: O(n)
	memory: O(1)
*/
func truncateSentence(s string, k int) string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
			if count == k {
				return s[:i]
			}
		}
	}
	return s
}
