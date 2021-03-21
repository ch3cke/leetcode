package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor_4(big int, medium int, small int) ParkingSystem {
	s := ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
	return s
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big > 0 {
			this.big -= 1
			return true
		} else {
			return false
		}
	case 2:
		if this.medium > 0 {
			this.medium -= 1
			return true
		} else {
			return false
		}
	case 3:
		if this.small > 0 {
			this.small -= 1
			return true
		} else {
			return false
		}
	default:
		return false
	}
}
