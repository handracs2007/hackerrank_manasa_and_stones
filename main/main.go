// https://www.hackerrank.com/challenges/manasa-and-stones/problem

package main

import (
	"fmt"
	"math"
)

type Int32Set struct {
	elements []int32
}

func NewInt32Set() *Int32Set {
	var int32Set = new(Int32Set)
	int32Set.elements = make([]int32, 0)

	return int32Set
}

func (i32Set *Int32Set) Add(element int32) {
	var exists = false
	for i := 0; i < len(i32Set.elements); i++ {
		if i32Set.elements[i] == element {
			exists = true
			break
		}
	}

	if !exists {
		i32Set.elements = append(i32Set.elements, element)
	}
}

func stones(n int32, a int32, b int32) []int32 {
	var result = NewInt32Set()

	var min = int32(math.Min(float64(a), float64(b)))
	var max = int32(math.Max(float64(a), float64(b)))
	var start = min * (n - 1)
	var end = max * (n - 1)

	for start <= end {
		result.Add(start)
		start += max - min

		if max == min {
			break
		}
	}

	return result.elements
}

func main() {
	fmt.Println(stones(3, 1, 2))
	fmt.Println(stones(4, 10, 100))
}
