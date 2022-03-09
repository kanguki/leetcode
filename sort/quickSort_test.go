package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func quickSort(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	//pick a pivot (l-1), place pivot in the right order. default is index 0
	//traverse array, if val < pivot, swap smaller value with current index of pivot, index of pivot++
	pivot := partition(arr)
	quickSort(arr[:pivot])
	quickSort(arr[pivot+1 : l])
}

func partition(arr []int) int {
	l := len(arr)
	pivot := arr[l-1]
	rightPosition := 0
	for i := 0; i < l-1; i++ {
		if arr[i] < pivot {
			swap(arr, i, rightPosition)
			rightPosition++
		}
	}
	swap(arr, rightPosition, l-1)
	return rightPosition
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func TestQuickSort(t *testing.T) {
	arr := []int{9, 4, 10, 3, 5, 7, 4}
	quickSort(arr)
	assert.Equal(t, []int{3, 4, 4, 5, 7, 9, 10}, arr)
	fmt.Println(arr)
}
