package main

import (
	"fmt"
)

func main() {
	input := "(1+(2*3)+((8)/4))+1"
	output := 3
	if maxDepth(input) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	time: O(n)
	memory: O(1)
*/
func maxDepth(s string) int {
	res := 0
	count := 0
	for _, c := range s {
		switch c {
		case '(':
			count++
		case ')':
			{
				if count <= 0 {
					return -1 // 题设说不出现，但仍需要保留
				} else {
					if count > res {
						res = count
					}
					count--
				}
			}
		default:
			continue
		}
	}
	return res
}
