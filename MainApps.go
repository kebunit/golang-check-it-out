package main

import (
	"fmt"
	// "github.com/kebunit/golang-check-it-out/basic"
	// "github.com/kebunit/golang-check-it-out/oop"
	// "github.com/kebunit/golang-check-it-out/network"
	"github.com/kebunit/golang-check-it-out/file"
	// "github.com/kebunit/golang-check-it-out/db"
)

func main() {
	// BASIC PACKAGE TEST
	// basic.CallGreeting()
	// basic.Hello()

	// OOP PACKAGE TEST
	// mycar := oop.NewCar(3242, 0, 23, 3422.34)
	// mycar.Print()
	// mycar.SetGasPedal(50000)
	// mycar.Print()
	// oop.TestOOP()

	// DB PACKAGE TEST
	// mydb.MyDBTest()

	// FILE PACKAGE 
	file.SaveToFile("filename.xbt", "Ternyata!!! as simple as that!")
	fmt.Println(file.ReadFile("filename.xbt"))
	// NETWORK PACKAGER TEST
	// network.ServerON()
	
}