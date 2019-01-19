// bubble_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func BubbleSort(arr []int) {
	for i := 0; i < len(arr) - 1; i ++ {
		for j := 0; j < len(arr) - i - 1; j ++ {
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
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
	BubbleSort(arr)
	fmt.Println("after: ", arr)
}