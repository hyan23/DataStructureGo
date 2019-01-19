// qsort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func QSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	low, high := 0, len(arr) - 1
	pivot := arr[low]
	arr[high], arr[low] = arr[low], arr[high]
	mid := low
	for i := low; i < high; i ++ {
		if arr[i] < pivot {
			arr[mid], arr[i] = arr[i], arr[mid]
			mid ++
		}
	}
	arr[mid], arr[high] = arr[high], arr[mid]
	QSort(arr[low:mid])
	QSort(arr[mid+1:high+1])
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
	QSort(arr)
	fmt.Println("after: ", arr)
}