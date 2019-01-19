// binary_tree_sort.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
)

type IncArray struct {
	arr *[]int
	i int
}

func (incArray *IncArray) Insert (value int) {
	(*incArray.arr)[incArray.i] = value
	incArray.i ++
}

type Node struct {
	value int
	left *Node
	right *Node
}

func (root *Node) Init() {
	root.left = root
}

func (root *Node) Insert (value int) {
	if root.left == root {
		root.value = value
		root.left = nil
	} else {
		if value < root.value {
			if root.left == nil {
				root.left = &Node{value, nil, nil}
			} else {
				root.left.Insert(value)
			}
		} else {
			if root.right == nil {
				root.right = &Node{value, nil, nil}
			} else {
				root.right.Insert(value)
			}
		}
	}
}

func (root *Node) InOrder (incArray *IncArray) {
	if root == nil {
		return
	}
	root.left.InOrder(incArray)
	incArray.Insert(root.value)
	root.right.InOrder(incArray)
}

func BinaryTreeSort(arr []int) {
	var root Node
	root.Init()
	for _, e := range arr {
		root.Insert(e)
	}
	incArray := IncArray{ &arr, 0 }
	root.InOrder(&incArray)
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
	BinaryTreeSort(arr)
	fmt.Println("after: ", arr)
}