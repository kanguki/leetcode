package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(canContruct2("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canContruct2("abcde", []string{"ab", "abc", "cd", "def", "abcd"}))
	fmt.Println(canContruct("abcdeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", []string{"ab", "abc", "cd", "def", "abcd"}))
	//fmt.Println(canContruct("abcdefabcd", []string{"ab", "abc", "cd", "def", "abcd"}))
}

//consume less space, but more time
func canContruct2(target string, s []string) bool {
	fmt.Println(target, s)
	tab := make([]bool, len(target)+1) //  //tab[i] represents answer for target[:i] (i chracters)
	tab[0] = true
	for i := 0; i < len(tab); i++ {
		if tab[i] {
			for _, v := range s {
				if i+len(v) < len(tab) && strings.HasPrefix(target, string(target[:i])+v) {
					tab[i+len(v)] = true
				}
			}
		}
	}
	for i := range tab {
		fmt.Println(i, target[:i], tab[i])
	}
	return tab[len(target)]
}

//idea:
//target: abcd; s: [a ab c cd bcd]
//=> tab: "" a ab abc abcd
func canContruct(target string, s []string) bool {
	tab := make([]string, len(target)+1)
	ok := make(map[string]bool)
	for i := 0; i < len(tab); i++ {
		tab[i] = string(target[0:i])
		ok[tab[i]] = false
	}
	ok[""] = true
	for i := 0; i < len(tab); i++ {
		if ok[tab[i]] {
			for _, v := range s {
				if _, exist := ok[tab[i]+string(v)]; exist {
					ok[tab[i]+string(v)] = true
				}
			}
		}
	}
	fmt.Println(tab, ok)
	return ok[target]
}
