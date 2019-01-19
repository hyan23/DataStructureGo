// heap_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func max_heapify(arr []int, start int, end int) {
	parent := start
	child := parent * 2 + 1
	for child <= end {
		if child + 1 <= end {
			if arr[child + 1] > arr[child] {
				child = child + 1
			}
		}
		if arr[parent] > arr[child] {
			return
		} else {
			arr[parent], arr[child] = arr[child], arr[parent]
			parent = child
			child = parent * 2 + 1
		}
	}
}

func HeapSort(arr []int) {
	end := len(arr) - 1
	for i := (end - 1) / 2; i >= 0; i -- {
		max_heapify(arr, i, end)
	}
	for i := end; i > 0; i -- {
		arr[i], arr[0] = arr[0], arr[i]
		max_heapify(arr, 0, i - 1)
	}
}

func RandomArray(len int, lb int, ub int) []int {
	arr := make([]int, len)
	for i, _ := range arr {
		arr[i] = lb + rand.Intn(ub)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().Unix())
	arr := RandomArray(10, 1, 10)
	fmt.Println("before: ", arr)
	HeapSort(arr)
	fmt.Println("after: ", arr)
}