package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// INT
	var intNum int16 = 32767
	// nếu để type là int16, thì +1 ở đây sẽ ra giá trị k xác định
	intNum = intNum + 1 // -32768
	fmt.Println(intNum)

	// FLOAT
	var floatNum = 12345678.9
	fmt.Println(floatNum)

	// CASTING
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result = floatNum32 + float32(intNum32)
	fmt.Println(result) // 12.1

	// STRING
	var myString = "Hellô" + " " + "World"
	fmt.Println(myString)
	fmt.Println(len(myString))
	fmt.Println(utf8.RuneCountInString(myString))

	// RUNE
	var myRune = 'a'
	fmt.Println(myRune)

	const myConst string = "Hello World"
	fmt.Println(myConst)
}
