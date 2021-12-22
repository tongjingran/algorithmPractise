package main

import (
	"fmt"
	"sort"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	input := struct {
		houses  []int
		heaters []int
	}{[]int{1, 2, 3}, []int{2}}
	output := 1
	if findRadius(input.houses, input.heaters) != output {
		panic("wrong answer")
	}
	fmt.Println("pass")
	return
}

/*
	find min radius
	time: O((m+n)logn)
*/
func findRadius(houses []int, heaters []int) int {
	max := 0
	sort.Ints(houses)
	sort.Ints(heaters)
	for _, house := range houses {
		min := abs(house, heaters[0])
		for _, heater := range heaters {
			if abs(house, heater) > min {
				break
			}
			min = abs(house, heater)
		}
		if min > max {
			max = min
		}
	}
	return max
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
