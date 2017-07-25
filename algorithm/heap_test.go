package algorithm

import (
	"testing"
)

type Array []int

func (arr Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}

func (arr Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func (arr Array) Len() int {
	return len(arr)
}

func TestHeap(t *testing.T) {
	var arr Array = Array{4, 21, 3, 1, 45, 632, 32, 1, 4}
	CreateHeap(arr)
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
}
