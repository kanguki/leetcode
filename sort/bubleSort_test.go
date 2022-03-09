package sort

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func bubleSort(arr []int) {
	//put the largest value in it right order
	l := len(arr)
	for loop := 0; loop < l-1; loop++ {
		//when there is no swap, we can stop here
		firstPartIsOrdered := true
		for i := 0; i < l-1-loop; i++ {
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
				firstPartIsOrdered = false
			}
		}
		if firstPartIsOrdered {
			break
		}
	}
}

func TestBubleSort(t *testing.T) {
	arr := []int{9, 4, 10, 3, 5, 7, 4}
	bubleSort(arr)
	assert.Equal(t, []int{3, 4, 4, 5, 7, 9, 10}, arr)
	fmt.Println(arr)
}
