package main

import "fmt"

func main() {
	fmt.Println(howSum(300, []int{100, 7, 14, 30}))
}

//return any subarray whose sum = target
func howSum(target int, nums []int) []int {
	tab := make([][]int, target+1)
	tab[0] = []int{}
	for i := 0; i < target+1; i++ {
		if tab[i] != nil {
			for _, v := range nums {
				if i+v < target+1 {
					copied := make([]int, len(tab[i]))
					copy(copied, tab[i])
					copied = append(copied, v)
					tab[i+v] = copied
				}
			}
		}
	}
	for i := range tab {
		fmt.Println(i, tab[i])
	}
	return tab[target]

}
