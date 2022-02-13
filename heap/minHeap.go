package main

import "fmt"

type MinHeap struct {
	arr []int
}

func left(index int) int {
	return index*2 + 1
}

func right(index int) int {
	return index*2 + 2
}

func parent(index int) int {
	return (index - 1) / 2
}

func (m *MinHeap) swap(index1, index2 int) {
	m.arr[index1], m.arr[index2] = m.arr[index2], m.arr[index1]
}

func (m *MinHeap) insert(val int) {
	m.arr = append(m.arr, val)
	m.heapifyUp(len(m.arr) - 1)
}

//place last element in the right order (parent <= child)
func (m *MinHeap) heapifyUp(index int) {
	for m.arr[parent(index)] > m.arr[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

func (m *MinHeap) pop() (poppedValue int) {
	l := len(m.arr)
	if l == 0 {
		return -1
	}
	poppedValue = m.arr[0]
	//put last-element's value to head, slice last element out
	m.arr[0] = m.arr[l-1]
	m.arr = m.arr[:l-1]
	m.heapifyDown(0)
	return poppedValue
}

//place first element in the right order
func (m *MinHeap) heapifyDown(index int) {
	last := len(m.arr) - 1
	childToSwap := 0
	l := left(index)
	r := right(index)
	for l <= last { //while it still has child(ren)
		if l == last { //only as one child
			childToSwap = l
		} else { //has 2 children, pick the one with smaller value

			if m.arr[l] < m.arr[r] {
				childToSwap = l
			} else {
				childToSwap = r
			}

		}

		if m.arr[index] > m.arr[childToSwap] {
			m.swap(index, childToSwap)
			index = childToSwap
			l = left(index)
			r = right(index)
		} else { //it's in the right order
			return
		}
	}
}

func main() {
	heap := MinHeap{arr: []int{}}
	arr := []int{2, 3, 8, 1, 5, 7, 0, 9, 6, 10, 4}
	for _, v := range arr {
		heap.insert(v)
	}
	fmt.Println(heap.arr)
	heap.pop()
	fmt.Println(heap.arr)
	heap.pop()
	fmt.Println(heap.arr)
	for range heap.arr {
		fmt.Print(heap.pop())
	}
}
