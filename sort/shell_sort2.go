// shell_sort2.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func ShellSort(arr []int) {
	for inc := len(arr) / 2; inc >= 1; inc /= 2 {
		for i := 0; i < len(arr) - inc; i ++ {
			for j := i + inc; j < len(arr); j += inc {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
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
	ShellSort(arr)
	fmt.Println("after: ", arr)
}