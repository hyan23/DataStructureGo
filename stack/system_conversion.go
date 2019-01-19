// system_conversion.go
// Author: hyan23
// Date: 2018.11.01

package main

import (
	"fmt"
	"strings"
)

func convert(base uint, number uint) string {
	var stack Stack
	stack.Init(64)
	for number > 0 {
		stack.Push(number % base)
		number /= base
	}
	result := make([]string, stack.Size())
	for i := 0; !stack.Empty(); i ++ {
		e, _ := stack.Pop()
		result[i] = fmt.Sprintf("%v", e)
	}
	return strings.Join(result, "")
}

func main() {
	fmt.Println(convert(8, 1348))
}