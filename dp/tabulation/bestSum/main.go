package main

import "fmt"

func main() {
	fmt.Println(bestSum(300, []int{7, 14}))
}
func bestSum(target int, nums []int) []int {
	tab := make([][]int, target+1) //tab[i] represents answer for target[:i] (i chracters)
	tab[0] = []int{}
	for i := 0; i < len(tab); i++ {
		if tab[i] != nil {
			for _, v := range nums {
				if i+v < len(tab) {
					copied := make([]int, len(tab[i]))
					copy(copied, tab[i])
					copied = append(copied, v)
					if tab[i+v] == nil {
						tab[i+v] = copied
					} else if len(copied) < len(tab[i+v]) {
						tab[i+v] = copied
					}
				}
			}
		}
	}
	for i := range tab {
		fmt.Println(i, tab[i])
	}
	return tab[target]
}
