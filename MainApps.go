package main

import (
	// "github.com/kebunit/golang-check-it-out/basic"
	"github.com/kebunit/golang-check-it-out/oop"
	// "github.com/kebunit/golang-check-it-out/db"
)

func main() {
	// BASIC PACKAGE TEST
	// basic.CallGreeting()
	// basic.Hello()
	// basic.Goserver()

	// OOP PACKAGE TEST
	mycar := oop.NewCar(3242, 0, 23, 3422.34)
	mycar.Print()
	mycar.SetGasPedal(50000)
	mycar.Print()
	oop.TestOOP()

	// DB PACKAGE TEST
	// mydb.MyDBTest()

	
}