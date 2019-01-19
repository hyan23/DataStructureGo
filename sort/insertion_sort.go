// insertion_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i ++ {
		key := arr[i]
		j := i - 1
		for ; j >= 0; j -- {
			if arr[j] > key {
				arr[j + 1] = arr[j]
			} else {
				break
			}
		}
		arr[j + 1] = key
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
	InsertionSort(arr)
	fmt.Println("after: ", arr)
}