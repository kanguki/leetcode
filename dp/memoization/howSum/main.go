package main

import "fmt"

func main() {
	//fmt.Println(allPossibleSubArraySum(7, []int{5,3,4,7}))
	fmt.Println(findLength([]int{1, 0, 1, 0, 0, 0, 0, 0, 1, 1}, []int{1, 1, 0, 1, 1, 0, 0, 0, 0, 0}))
}
func findLength(nums1 []int, nums2 []int) int {
	maxSoFar := 0
	left, right := 0, 1
	for right <= len(nums1) {
		for right <= len(nums1) && isSubarray(nums1[left:right], nums2) {
			fmt.Println(nums1[left:right])
			maxSoFar = max(maxSoFar, right-left)
			right++
		}
		left, right = right, right+1
	}
	return maxSoFar
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isSubarray(shorter, longer []int) bool {
	if len(shorter) > len(longer) || len(shorter)*len(longer) == 0 {
		return false
	}
	l := 0
	for l < len(longer) && len(longer)-l >= len(shorter) {
		s := 0
		curL := l
		for s < len(shorter) && l < len(longer) && shorter[s] == longer[l] {
			s++
			l++
		}
		if s == len(shorter) {
			return true
		}
		l = curL + 1
	}
	return false
}
