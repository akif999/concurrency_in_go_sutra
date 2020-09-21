package main

import "fmt"

func main() {
	multiply := func(value, multiplier int) int {
		return value * multiplier
	}

	add := func(value, additive int) int {
		return value + additive
	}

	ints := []int{1, 2, 3, 4}
	for _, v := range ints {
		fmt.Println(multiply(add(multiply(v, 2), 1), 2))
	}
}

// func main() {
// 	multiply := func(values []int, multiplier int) []int {
// 		multipiedValues := make([]int, len(values))
// 		for i, v := range values {
// 			multipiedValues[i] = v * multiplier
// 		}
// 		return multipiedValues
// 	}
//
// 	add := func(values []int, additive int) []int {
// 		addedValues := make([]int, len(values))
// 		for i, v := range values {
// 			addedValues[i] = v + additive
// 		}
// 		return addedValues
// 	}
//
// 	ints := []int{1, 2, 3, 4}
// 	for _, v := range multiply(add(multiply(ints, 2), 1), 2) {
// 		fmt.Println(v)
// 	}
// }

// func main() {
// 	multiply := func(values []int, multiplier int) []int {
// 		multipiedValues := make([]int, len(values))
// 		for i, v := range values {
// 			multipiedValues[i] = v * multiplier
// 		}
// 		return multipiedValues
// 	}
//
// 	add := func(values []int, additive int) []int {
// 		addedValues := make([]int, len(values))
// 		for i, v := range values {
// 			addedValues[i] = v + additive
// 		}
// 		return addedValues
// 	}
//
// 	ints := []int{1, 2, 3, 4}
// 	for _, v := range add(multiply(ints, 2), 1) {
// 		fmt.Println(v)
// 	}
// }
