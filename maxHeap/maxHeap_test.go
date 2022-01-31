package maxHeap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertHeap(t *testing.T) {
	testInput := []int{10, 20, 30}
	m := &MaxHeap{}
	for _, v := range testInput {
		m.Insert(v)
	}
	assert.EqualValuesf(t, m.array, []int{30, 10, 20}, "30,10,20 should be the heapified value")
}

func TestPopHeap(t *testing.T) {
	testInput := []int{10, 20, 30, 5, 7, 9}
	m := &MaxHeap{}
	for _, v := range testInput {
		m.Insert(v)
	}

	res := make([]int, 0, len(testInput))
	for i := 0; i < len(testInput); i++ {
		v := m.Pop()
		res = append(res, v)
	}
	assert.EqualValuesf(t, res, []int{30, 20, 10, 9, 7, 5}, "should pop largest element in heap")
}
