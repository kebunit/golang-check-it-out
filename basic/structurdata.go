package basic

import "fmt"

const usixteenbitmax float64 = 66555
const kmh_multiple float64 = 1.60606

type car struct {
	gas_pedal uint16
	brake_pedal uint16
	steering_wheel int16
	top_speed_kwh float64
}

func (c car) kmh() float64{
	return float64(c.gas_pedal) * (c.top_speed_kwh/usixteenbitmax)
}

func (c car) mph() float64{
	return float64(c.gas_pedal) * (c.top_speed_kwh/usixteenbitmax/kmh_multiple)
}

func MobilanYuk() {
	a_car := car{ gas_pedal:23424,
				brake_pedal:0,
				steering_wheel:18388,
				top_speed_kwh:255.0}

	fmt.Println(a_car.gas_pedal)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}