// bucket_sort.go
// Author: hyan23
// Date: 2018.10.29

package main

import (
	"fmt"
	"time"
	"math/rand"
	"math"
	"container/list"
)

const BUCKETS = 10

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

func BucketSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	
	max, min := max(arr), min(arr)
	num := max - min + 1
	if num == 1 {
		return
	}
	
	buckets := make([]*list.List, BUCKETS)
	for i, _ := range buckets {
		buckets[i] = list.New()
		// buckets[i].Init()
	}
	
	interval := int (math.Ceil(float64 (num) / BUCKETS))
	for _, e := range arr {
		i := (e - min) / interval
		buckets[i].PushBack(e)
	}
	
	buckets2 := make([][]int, BUCKETS)
	for i, bucket := range buckets {
		buckets2[i] = make([]int, bucket.Len())
		for j, e := 0, bucket.Front(); e != nil; e = e.Next() {
			buckets2[i][j] = e.Value.(int)
			j ++
		}
		BucketSort(buckets2[i])
	}
	
	for i, pos := 0, 0; i < BUCKETS; i ++ {
		for _, e := range buckets2[i] {
			arr[pos] = e
			pos ++
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
	arr := RandomArray(32, 1, 999)
	fmt.Println("before: ", arr)
	BucketSort(arr)
	fmt.Println("after: ", arr)
}