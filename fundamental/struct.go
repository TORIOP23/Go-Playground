package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.mpg * e.gallons
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("You can't make it!")
	}
}

func main() {
	//var myEngine = gasEngine{25, 15}
	var myEngine = electricEngine{25, 15}
	//fmt.Println(myEngine.mpg, myEngine.gallons)
	canMakeIt(myEngine, 50)
}
