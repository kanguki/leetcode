package main

func main() {
	println(waysToTravelFromUpLeftToBottomRightOfGrid(18, 18))
}

//idea: plus right and down with current position
func waysToTravelFromUpLeftToBottomRightOfGrid(row, col int) int {
	tab := make([][]int, row) //store number of ways to go from up most left to down most right
	for i := range tab {
		tab[i] = make([]int, col)
	}
	tab[0][0] = 1 //there's one way to go to itself
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			current := tab[r][c] //number of ways to go to current will adds to number of ways to go to its right/ down node.
			if r+1 < row {
				tab[r+1][c] += current
			}
			if c+1 < col {
				tab[r][c+1] += current
			}
		}
	}
	return tab[row-1][col-1]
}
