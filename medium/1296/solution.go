package main

import (
	"fmt"
)

func main() {
	input := struct {
		nums []int
		k    int
	}{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4}
	output := true
	if isPossibleDivide(input.nums, input.k) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	right -> root -> left
	time: O(n*k)
	memory: O(n)
*/
func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k != 0 {
		return false
	}
	min := 1000000000
	max := 0
	for _, i := range nums {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}
	a := make([]int, max-min+1)
	for _, i := range nums {
		a[i-min]++
	}
	for i := 0; i <= len(a)-k; i++ {
		if a[i] == 0 {
			continue
		}
		for j := 1; j < k; j++ {
			if a[i+j] < 0 {
				return false
			}
			a[i+j] = a[i+j] - a[i]
		}
		a[i] = 0
	}
	for i := range a {
		if a[i] != 0 {
			return false
		}
	}
	return true
}
