package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mergeSort(arr []int) []int {
	l := len(arr)
	//pick the mid element, sort left and right, then merge 2 sorted array
	if l <= 1 {
		return arr
	}
	mid := l / 2
	left := arr[:mid]
	right := arr[mid:]
	left = mergeSort(left)
	right = mergeSort(right)
	return merge2SortedArrays(left, right)
}

func merge2SortedArrays(a1, a2 []int) []int {
	l1, l2 := len(a1), len(a2)
	smol1, smol2 := 0, 0
	sortedArray := []int{}
	for smol1 < l1 && smol2 < l2 {
		if a1[smol1] < a2[smol2] {
			sortedArray = append(sortedArray, a1[smol1])
			smol1++
		} else {
			sortedArray = append(sortedArray, a2[smol2])
			smol2++
		}
	}
	if smol1 == l1 {
		return append(sortedArray, a2[smol2:l2]...)
	}
	if smol2 == l2 {
		return append(sortedArray, a1[smol1:l1]...)
	}
	return sortedArray
}
func TestMergeSort(t *testing.T) {
	arr := []int{9, 4, 10, 3, 5, 7, 4}
	a := mergeSort(arr)
	assert.Equal(t, []int{3, 4, 4, 5, 7, 9, 10}, a)
	fmt.Println(a)
}
