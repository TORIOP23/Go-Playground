package main

import (
	"fmt"
)

func main() {
	intArr := [3]int{1, 2, 3}

	fmt.Println(intArr)

	intSlice := []int{4, 5, 6}
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// The length is 3 with capacity 3

	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// The length is 4 with capacity 6

	intSlice2 := []int{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))
	// The length is 6 with capacity 6
	// if intSlice2 = 8,9,10 then
	// -> The length is 7 with capacity 12
}
