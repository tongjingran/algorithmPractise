package main

import (
	"fmt"
	"strings"
)

func main() {
	input := struct {
		licensePlate string
		words        []string
	}{"1s3 456", []string{"looks", "pest", "stew", "show"}}
	output := "pest"
	if shortestCompletingWord(input.licensePlate, input.words) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	find shortest competing word
	time: O(n+m+sum(l))
	memory: O(sum(l))
*/
func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlate = strings.ToLower(licensePlate)
	var m1 = [26]int{}
	for _, ch := range licensePlate {
		if ch < 'a' || ch > 'z' {
			continue
		}
		m1[ch-'a']++
	}
	var res string
	for _, word := range words {
		if res != "" && len(word) >= len(res) {
			continue
		}
		m2 := [26]int{}
		for _, ch := range word {
			m2[ch-'a']++
		}
		tmp := true
		for i := range m1 {
			if m2[i] < m1[i] {
				tmp = false
				break
			}
		}
		if tmp {
			res = word
		}
	}
	return res
}
