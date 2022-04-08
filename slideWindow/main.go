package main

import "fmt"

func main() {
	//fmt.Println(minLengthSubarraySumGteK([]int{2, 3, 5, 4, 2, 6}, 9))
	// fmt.Println(longestSubStringOfKDistinctCharacters("acccpbbebiiiiii", 3))
	fmt.Println(minimumSizeSubArraySum(7, []int{2, 3, 1, 2, 4, 3}))
}

//leetcode 209
func minimumSizeSubArraySum(target int, nums []int) int {
	start, end, sum, minSoFar := 0, 0, 0, int(1e5)
	//for sum > target -> update sum, update start
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			minSoFar = min(minSoFar, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if minSoFar == 1e5 {
		minSoFar = 0
	}
	return minSoFar
}

//given a string, find longest substring that has k distinct characters. if it doesn't exit, return ""
func longestSubStringOfKDistinctCharacters(s string, maxDistinct int) string {
	res, m, maxSoFar, start, end := "", make(map[string]int), 0, 0, 0
	for end < len(s) {
		m[string(s[end])]++
		for len(m) > maxDistinct {
			cha := string(s[start])
			m[cha]--
			if m[cha] == 0 {
				delete(m, cha)
			}
			start++
		}
		if len(m) == maxDistinct && maxSoFar < end-start+1 {
			maxSoFar = end - start + 1
			res = string(s[start : end+1])
		}
		end++
		//		second approach, but it's not clear end will move forward anyway, so I don't recommend this.
		//		if len(m) <= maxDistinct {
		//			if len(m) == maxDistinct && maxSoFar < end - start +1 {
		//				maxSoFar = end - start +1
		//				res = string(s[start: end+1])
		//			}
		//			end++
		//		} else {
		//			cha := string(s[start])
		//			m[cha]--
		//			if m[cha] == 0 {
		//				delete(m ,cha)
		//			}
		//			start++
		//		}
	}
	return res
}

//find minimum contiguous subarray whose sum >= k
func minLengthSubarraySumGteK(arr []int, k int) []int {
	start, end, sum, res, minLenSoFar := 0, 0, 0, []int{}, int(1e5)
	for end < len(arr) {
		sum += arr[end]
		for sum >= k {
			if minLenSoFar > (end - start + 1) {
				res = arr[start : end+1]
				minLenSoFar = end - start + 1
			}
			sum -= arr[start]
			start++
		}
		end++
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
