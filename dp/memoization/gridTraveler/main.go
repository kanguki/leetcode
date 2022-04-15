/**
* number of ways to travel a 2D grid of size m * n (m rows, n cols)
* you may only move down or right
**/
package main

import "fmt"

func main() {
	// fmt.Println(fastTravel(100, 100))
	// fmt.Println(uniquePaths(23, 12))
	//fmt.Println(backtrack(50,50))
	//fmt.Println(recursive(50,50))
}

type coor struct {
	x, y int
}

func uniquePaths(m int, n int) int {
	uniquePathsOfGridMemo := make(map[string]int)
	var solve func(int, int) int
	solve = func(rows, cols int) int {
		key := fmt.Sprintf("%v_%v", rows, cols)
		if uniquePathsOfGridMemo[key] != 0 {
			return uniquePathsOfGridMemo[key]
		}
		if rows*cols == 1 {
			return 1
		} //we reach the target
		if rows*cols == 0 {
			return 0
		} //we're out of grid
		uniquePathsOfGridMemo[key] = uniquePaths(rows-1, cols) + uniquePaths(rows, cols-1)
		return uniquePathsOfGridMemo[key]
	}
	return solve(m, n)
}

//time: O(m*n) - we can form m*n grid that has size <= m*n
//space: O(m*n+m+n) - layers + memo
func fastTravel(m, n int) int {
	memo := make(map[string]int)
	var recursiveWithMemo func(int, int, map[string]int) int
	recursiveWithMemo = func(rows, cols int, memo map[string]int) int {
		//basically, paths to move in a m*n grid is the same as n*m grid, so we can check for memo[cols_rows] too.
		key := fmt.Sprintf("%v_%v", rows, cols)
		reversedKey := fmt.Sprintf("%v_%v", cols, rows)
		pathsOfGrid := memo[key]
		if pathsOfGrid == 0 {
			pathsOfGrid = memo[reversedKey]
		}
		if pathsOfGrid != 0 { //use memo if exist
			return pathsOfGrid
		}
		if rows*cols <= 1 {
			return rows * cols
		}
		memo[key] = recursiveWithMemo(rows-1, cols, memo) + recursiveWithMemo(rows, cols-1, memo)
		memo[reversedKey] = memo[key]
		return memo[key]
	}
	return recursiveWithMemo(m, n, memo)

}

func recursive(m, n int) int {
	if m*n <= 1 {
		return m * n
	}
	return recursive(m-1, n) + recursive(m, n-1)
}

//time: O(2^(m+n)) - start from (0,0), at each position it reaches, there are 2 ways of proceeding. max number of layers is m+n (|_), start from (0,0)
//space: O(m+n) - max number of layers
func backtrack(m, n int) int {
	ways := 0
	debugCount := 0
	var travel func(coor)
	travel = func(c coor) {
		//fmt.Println(c)
		debugCount++
		if c.x == m-1 && c.y == n-1 {
			ways++
			return
		}
		if isValid(coor{x: c.x + 1, y: c.y}, m, n) {
			travel(coor{x: c.x + 1, y: c.y})
		}
		if isValid(coor{x: c.x, y: c.y + 1}, m, n) {
			travel(coor{x: c.x, y: c.y + 1})
		}
	}
	travel(coor{0, 0})
	//fmt.Println("ops: ", debugCount)
	//fmt.Println("ways: ", ways)
	return ways
}

func isValid(c coor, m, n int) bool {
	return c.x < m && c.y < n
}
