package main

import "fmt"

func main() {
	fmt.Println(canSum(7, []int{2, 3}))
	fmt.Println(fastCanSum(7, []int{2, 3}, nil))
}
func fastCanSum(target int, nums []int, memo map[int]bool) bool {
	if memo == nil {
		memo = make(map[int]bool)
	}
	if memo[target] {
		return true
	}
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	for _, v := range nums {
		if fastCanSum(target-v, nums, memo) {
			memo[target] = true
			return true
		}
	}
	return false
}
func canSum(target int, nums []int) bool {
	if target == 0 {
		return true
	}
	if target < 0 {
		return false
	}
	for _, v := range nums {
		if canSum(target-v, nums) {
			return true
		}
	}
	return false
}
