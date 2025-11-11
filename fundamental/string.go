package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "Résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	//0 82
	//1 233
	//3 115
	//4 117
	//5 109
	//6 233

	// 8
	fmt.Printf("The length of 'myString is %v\n", len(myString))

	fmt.Println("-----rune-----")
	var myString2 = []rune("Résumé")
	var indexed2 = myString2[0]
	fmt.Printf("%v, %T\n", indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	fmt.Printf("String Builder")
	var strSlice = []string{"S", "u", "b", "s", "c"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Printf("\n%®v", strBuilder.String())
}
