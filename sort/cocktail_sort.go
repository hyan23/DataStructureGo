// cocktail_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func CocktailSort(arr []int) {
	low, high := 0, len(arr) - 1
	for ; low < high; {
		for i := low; i < high; i ++ {
			if arr[i] > arr[i + 1] {
				arr[i], arr[i + 1] = arr[i + 1], arr[i]
			}
		}
		high --
		for i := high; i > low; i -- {
			if arr[i] < arr[i - 1] {
				arr[i], arr[i - 1] = arr[i - 1], arr[i]
			}
		}
		low ++
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
	CocktailSort(arr)
	fmt.Println("after: ", arr)
}