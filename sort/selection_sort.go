// selection_sort.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func SelectionSort(arr []int) {
	for i := 0; i < len(arr); i ++ {
		min := i
		for j := i + 1; j < len(arr); j ++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
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
	SelectionSort(arr)
	fmt.Println("after: ", arr)
}