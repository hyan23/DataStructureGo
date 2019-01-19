// qsort2.go
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
	key := arr[0]
	low, high := 0, len(arr) - 1
	for i := 1; i <= high; {
		if arr[i] > key {
			arr[high], arr[i] = arr[i], arr[high]
			high --
		} else {
			arr[low], arr[i] = arr[i], arr[low]
			low ++
			i ++
		}
	}
	QSort(arr[:low])
	QSort(arr[low+1:])
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