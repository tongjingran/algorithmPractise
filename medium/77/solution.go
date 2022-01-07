package main

import (
	"fmt"
)

func main() {
	input := struct {
		n int
		k int
	}{4, 2}
	fmt.Println(combine(input.n, input.k))
	return
}

var res [][]int

/*
	all combination
	time: O(n*2^n)
	memory: O(n)
*/
func combine(n int, k int) [][]int {
	res = [][]int{}
	cb(1, n, k, []int{})
	return res
}

/*
   对于每一个位置，只有选中或者不选中
   若选中，则后续为combine([i+1:n],k-1)
   若不选中，则后续为combine([i+1:n],k)
   若n-i==k，则返回后续所有值[i+1:n]
*/
func cb(start, end, k int, arr []int) {
	if k == 0 {
		res = append(res, arr)
		return
	}
	if end-start+1 == k {
		for i := start; i <= end; i++ {
			arr = append(arr, i)
		}
		res = append(res, arr)
		return
	}
	// 选中start
	selectArr := []int{}
	for _, i := range arr {
		selectArr = append(selectArr, i)
	}
	selectArr = append(selectArr, start)
	cb(start+1, end, k-1, selectArr)
	cb(start+1, end, k, arr)
	return
}
