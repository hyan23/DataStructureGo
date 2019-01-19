// counting_sort.go
// Author: hyan23
// Date: 2018.10.30

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func min(arr[] int) int {
	m := arr[0]
	for _, e := range arr {
		if e < m {
			m = e
		}
	}
	return m
}

func max(arr[] int) int {
	m := arr[0]
	for _, e := range arr {
		if e > m {
			m = e
		}
	}
	return m
}

func CountingSort(arr []int) {
	min, max := min(arr), max(arr)
	C := make([]int, max - min + 1)
	for _, e := range arr {
		C[e - min] ++
	}
	for i := 1; i < len(C); i ++ {
		C[i] += C[i - 1]
	}
	temp := make([]int, len(arr))
	copy(temp[:], arr[:])
	for _, e := range temp {
		arr[C[e - min] - 1] = e
		C[e - min] --
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
	CountingSort(arr)
	fmt.Println("after: ", arr)
}