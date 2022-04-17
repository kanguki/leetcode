package main

import "fmt"

func main() {
	println(canSum(300, []int{7, 14}))
}
func canSum(target int, nums []int) bool {
	tab := make([]bool, target+1) // //tab[i] represents answer for target[:i] (i chracters)
	tab[0] = true
	for i := 0; i < target; i++ {
		if tab[i] {
			for _, v := range nums {
				if i+v < target+1 { //still in bound
					tab[i+v] = true
				}
			}
		}
	}
	fmt.Println(tab)
	return tab[target]
}
