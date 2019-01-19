// radix_sort.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
	"container/list"
	"math"
)

func RadixSort(arr []int) {
	buckets := make([]*list.List, 10)
	for i := 0; i < 10; i ++ {
		buckets[i] = list.New()
	}
	max_divisor := int (math.Pow10(int (math.Log10(math.MaxInt64))))
	// 只能处理位于[0,max_divisor-1]之间的数据
	for i := 10; i <= max_divisor; i *= 10 {
		for _, e := range arr {
			buckets[e % i / (i / 10)].PushBack(e)
		}
		pos := 0
		for _, l := range buckets {
			for e := l.Front(); e != nil; e = e.Next() {
				arr[pos] = e.Value.(int)
				pos ++
			}
		}
		for i := 0; i < 10; i ++ {
			buckets[i].Init()
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
	arr := RandomArray(10, 1, 1000)
	fmt.Println("before: ", arr)
	RadixSort(arr)
	fmt.Println("after: ", arr)
}