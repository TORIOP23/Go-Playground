package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]uint)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{
		"Adam":  23,
		"Sarah": 45,
		"Danny": 12,
	}
	fmt.Println(myMap2)

	var age, ok = myMap2["Sarah"]
	if ok {
		fmt.Println("Sarah is", age)
	} else {
		fmt.Println("Sarah is not found")
	}

	delete(myMap2, "Sarah")

	for name := range myMap2 {
		fmt.Printf("The name is %v\n", name)
	}
}
