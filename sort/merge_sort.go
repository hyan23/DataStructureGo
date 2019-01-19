// merge_sort.go
// Author: hyan23
// Date: 2018.10.30

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	left, right := arr[:len(arr)/2], arr[len(arr)/2:]
	MergeSort(left)
	MergeSort(right)
	temp := make([]int, len(arr))
	i, j, k := 0, 0, 0
	for ; i < len(left) && j < len(right); {
		if left[i] < right[j] {
			temp[k] = left[i]
			i ++
			k ++
		} else if left[i] > right[j] {
			temp[k] = right[j]
			j ++
			k ++
		} else {
			temp[k] = left[i]
			i ++
			k ++
			temp[k] = right[j]
			j ++
			k ++
		}
	}
	for ; i < len(left); {
		temp[k] = left[i]
		i ++
		k ++
	}
	for ; j < len(right); {
		temp[k] = right[j]
		j ++
		k ++
	}
	copy(arr[:], temp[:])
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
	MergeSort(arr)
	fmt.Println("after: ", arr)
}