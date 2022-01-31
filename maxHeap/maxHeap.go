package maxHeap

//https://youtu.be/3DYIgTC4T1o

import (
	"fmt"
	"sync"
)

type MaxHeap struct {
	array []int //slice that hold the array
	mx    sync.Mutex
}

func (h *MaxHeap) Insert(key int) {
	h.mx.Lock()
	defer h.mx.Unlock()
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1) // re-arrange
}

func (h *MaxHeap) Pop() int {
	h.mx.Lock()
	defer h.mx.Unlock()
	length := len(h.array)
	if length == 0 {
		fmt.Println("cannot pop since length is 0")
		return -1
	}

	first := h.array[0]
	h.array[0] = h.array[length-1]
	h.array = h.array[:length-1] // cut the last el
	h.maxHeapifyDown(0)
	return first
}

//bubble the node up
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

//heapify from top to down
func (h *MaxHeap) maxHeapifyDown(index int) {
	length := len(h.array)
	l, r := left(index), right(index)
	childToCompare := 0
	for l <= length-1 {
		if l == length-1 {
			//left child is the only child to compare
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			// left is larger one
			childToCompare = l
		} else {
			// right is larger one
			childToCompare = r
		}
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

//get the parent index
func parent(index int) int {
	return (index - 1) / 2
}

func left(parentIndex int) int {
	return 2*parentIndex + 1
}

func right(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}
