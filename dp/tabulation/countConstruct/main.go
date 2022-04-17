package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
}

//space: O(target), time: O(target^2*n)
func countConstruct(target string, words []string) int {
	tab := make([]int, len(target)+1) //tab[i] tells if target[:i] is contructable.  //tab[i] represents answer for target[:i] (i chracters)
	tab[0] = 1
	for i := range tab {
		if tab[i] > 0 {
			for _, word := range words {
				if strings.HasPrefix(target, string(target[:i])+word) {
					tab[i+len(word)] += tab[i]
				}
			}
		}
	}
	for i := range tab {
		fmt.Println(i, target[:i], tab[i])
	}
	return tab[len(target)]
}
