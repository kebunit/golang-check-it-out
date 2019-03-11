package test

import (
	"testing"
	"github.com/kebunit/golang-check-it-out/oop"
)

func TestOps(t * testing.T) {
	mycar := oop.NewCar(324, 182, 10, 100)

	if mycar.GetGasPedal() != 10 {
		t.Errorf("Tidak sama dengan 10, tetapi %d", mycar.GetGasPedal())
	} else {
		t.Errorf("sama dengan 10000000")
	}
 }