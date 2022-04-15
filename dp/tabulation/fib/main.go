package main

import "fmt"

func main() {
	fmt.Println(fib(50))
}

func fib(n int) int {
	tab := make([]int, n+1)
	tab[1] = 1
	for i := 2; i <= n; i++ {
		tab[i] = tab[i-1] + tab[i-2]
	}
	return tab[n]
}
