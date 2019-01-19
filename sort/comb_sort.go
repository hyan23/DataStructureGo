// comb_sort.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func RadixSort(arr []int) {
	for gap, swaps := len(arr), 0; gap > 1 || swaps > 0; {
		if gap > 1 {
			gap = int (float64 (gap) * 0.8)
		}
		swaps = 0
		for i := 0; i < len(arr) - gap; i ++ {
			if arr[i] > arr[i + gap] {
				arr[i], arr[i + gap] = arr[i + gap], arr[i]
				swaps ++
			}
		}
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
	RadixSort(arr)
	fmt.Println("after: ", arr)
}