package main

import "log"

func twoSumMap(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		diff := target - v
		if indx, ok := m[diff]; ok {
			return []int{indx, i}
		}
		m[v] = i
	}
	return nil
}

func twoSumOn2(nums []int, target int) []int {
	for i, v := range nums {
		for j, k := range nums {
			if i == j {
				continue
			}
			if v+k == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	target := 3
	log.Println(twoSumMap(a, target))
	log.Println(twoSumOn2(a, target))

}
