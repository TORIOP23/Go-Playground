package main

import "fmt"

func main() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
