package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(allConstruct("abcdef", []string{"ab", "abcd", "cd", "def", "abcd", "ef", "c"}))
}
func allConstruct(target string, words []string) [][]string {
	tab := make([][][]string, len(target)+1) //tab[i] represents answer for target[:i] (i chracters)
	tab[0] = [][]string{{}}                  //this initialize [][]
	for i := range tab {
		if tab[i] != nil {
			for _, alts := range tab[i] {
				if alts != nil {
					for _, word := range words {
						can := strings.Join(append(alts, word), "")
						if strings.HasPrefix(target, can) {
							copied := make([]string, len(alts))
							copy(copied, alts)
							copied = append(copied, word)
							if tab[len(can)] == nil {
								tab[len(can)] = [][]string{copied}
							} else {
								tab[len(can)] = append(tab[len(can)], copied)
							}
						}
					}
				}
			}
		}
	}
	for i := range tab {
		fmt.Println(i, target[:i], tab[i])
	}
	return tab[len(target)]
}
