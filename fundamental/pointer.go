package main

import "fmt"

func main() {
	var p = new(int32)
	var i int32
	p = &i
	*p = 20

	fmt.Printf("The value p point to is: %v\n", *p)
	fmt.Printf("The value i is: %v\n", i)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v\n", thing1)
}

// make a copy of parameter so we can use pointer here
func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] *= thing2[i]
	}

	return thing2
}
