package oop

import "fmt"

type Address struct {
	street string
	number int16
	rt int16
	rw int16
}

type Student struct {
	name string
	address Addresssss
}


func TestNested() {
	Mahasiswa := Student{
		name : "Sabituddin Bigbang",
		address : Address{
			street : "Jl. Medan Merdeka",
			number : 17,
			rt : 3,
			rw : 4,
		},
	}	

	fmt.Println(Mahasiswa)
}