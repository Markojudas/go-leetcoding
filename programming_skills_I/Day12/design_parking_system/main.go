package main

import "fmt"

type ParkingSystem struct {
	BigSpace   int
	MidSpace   int
	SmallSpace int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	ps := ParkingSystem{
		BigSpace:   big,
		MidSpace:   medium,
		SmallSpace: small,
	}

	return ps
}

func (ps *ParkingSystem) AddCar(carType int) bool {

	switch carType {
	case 1:
		if ps.BigSpace < 1 {
			return false
		} else {
			ps.BigSpace -= 1
		}
	case 2:
		if ps.MidSpace < 1 {
			return false
		} else {
			ps.MidSpace -= 1
		}
	case 3:
		if ps.SmallSpace < 1 {
			return false
		} else {
			ps.SmallSpace -= 1
		}
	default:
		return false
	}

	return true
}

func main() {

	obj := Constructor(1, 1, 0)

	fmt.Println(obj.AddCar(1)) //true
	fmt.Println(obj.AddCar(2)) //true
	fmt.Println(obj.AddCar(3)) //false
	fmt.Println(obj.AddCar(1)) //false

}
