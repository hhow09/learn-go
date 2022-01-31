package maxHeap

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func buildTestCase(n int) []int {
	res := make([]int, 0, n)
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		res = append(res, rand.Intn(2000))
	}
	return res
}

func copyAndSort(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)
	sort.Slice(output, func(i, j int) bool {
		return output[i] > output[j]
	})
	return output
}

func TestInsertHeap(t *testing.T) {
	testInput := []int{10, 20, 30}
	m := &MaxHeap{}
	for _, v := range testInput {
		m.Insert(v)
	}
	assert.EqualValuesf(t, m.array, []int{30, 10, 20}, "30,10,20 should be the heapified value")
}

func TestPopHeap(t *testing.T) {
	testInput := buildTestCase(200)
	m := &MaxHeap{}
	for _, v := range testInput {
		m.Insert(v)
	}

	res := make([]int, 0, len(testInput))
	for i := 0; i < len(testInput); i++ {
		v := m.Pop()
		res = append(res, v)
	}
	expected := copyAndSort(testInput)
	assert.EqualValuesf(t, res, expected, "should pop largest element in heap")
}

func TestPopHeapConcurrent(t *testing.T) {
	testInput := buildTestCase(5)
	m := &MaxHeap{}
	for _, v := range testInput {
		m.Insert(v)
	}

	res := make([]int, 0, len(testInput))
	c := make(chan int) // unbuffered channel
	defer close(c)
	for i := 0; i < len(testInput); i++ {
		go func() {
			v := m.Pop()
			c <- v
		}()
	}

	for i := 0; i < len(testInput); i++ {
		res = append(res, <-c)
	}
	expected := copyAndSort(testInput)
	assert.EqualValuesf(t, res, expected, "should pop largest element in heap")
}
