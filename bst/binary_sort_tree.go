// binary_sort_tree.go
// Author: hyan23
// Date: 2018.10.31

package main

import (
	"fmt"
	"time"
	"math/rand"
)

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

func (root *Node) InOrder () {
	if root == nil {
		return
	}
	root.left.InOrder()
	fmt.Println(root.value)
	root.right.InOrder()
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
	var root Node
	root.Init()
	for _, e := range arr {
		root.Insert(e)
	}
	fmt.Println("after: ")
	root.InOrder()
}