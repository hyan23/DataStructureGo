// quick_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func partition(arr []int, low int, high int) int {
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
	return mid
}

func quick_sort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := partition(arr, low, high)
	quick_sort(arr, low, mid - 1)
	quick_sort(arr, mid + 1, high)
}

func QSort(arr []int) {
	quick_sort(arr, 0, len(arr) - 1)
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