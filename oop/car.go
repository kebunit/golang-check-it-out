package oop

import "fmt"

// global package variable
const usixteenbitmax float64 = 66555
const kmh_multiple float64 = 1.60606

// model
type car struct {
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kwh float64
}

// create a new car
func NewCar(gp uint16, bp uint16, sw int16, tsk float64) car {
	return car{gp, bp, sw, tsk}
}

// getter
func (c car) GetGasPedal() uint16 { 
	return c.gas_pedal
}

func (c car) GetBrakePedal() uint16 {
	return c.brake_pedal
}

func (c car) GetSteeringWheel() int16 {
	return c.steering_wheel
}

func (c car) GetTopSpeedKwh() float64 {
	return c.top_speed_kwh
}

// setter
func (c * car) SetGasPedal(gp uint16) {
	c.gas_pedal = gp
}

func (c * car) SetBrakePedal(bp uint16) {
	c.brake_pedal = bp
}

func (c * car) SetSteeringWheel(sw int16) {
	c.steering_wheel = sw
}

func (c * car) SetTopSpeedKwh(tsk float64) {
	c.top_speed_kwh = tsk
}

// methods
func (c car) Kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kwh/usixteenbitmax)
}

func (c car) Mph() float64{
	return float64(c.gas_pedal) * (c.top_speed_kwh/usixteenbitmax/kmh_multiple)
}

func (c car) Print() {
	fmt.Println("------------- print a car ------------");
	fmt.Println("GAS PEDAL      : ", c.gas_pedal);
	fmt.Println("BRAKE PEDAL    : ", c.brake_pedal);
	fmt.Println("STEERING WHEEL : ", c.steering_wheel);
	fmt.Println("TOP SPEED KWH  : ",c.top_speed_kwh);
}

func TestOOP() {
	a_car := NewCar(23424, 0, 18388, 255.5)	
	a_car.Print()
	a_car.SetGasPedal(5000)
	a_car.SetTopSpeedKwh(563.234)
	a_car.Print()
	fmt.Println(a_car.GetGasPedal())
}

/*
NOTES :
you can use initial uppercase letter for other package use (public)
and then you can use initial lowercase letter for internal pacakge use (private)
*/