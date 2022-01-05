package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}
// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	// panic("Please implement the NewCar function")

	return Car {
		speed: speed,
		battery: 100,
		batteryDrain: batteryDrain,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	// panic("Please implement the NewTrack function")

	return Track {
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	// panic("Please implement the Drive function")
	if car.battery < car.batteryDrain {
		return car
	} else {
		car.distance += car.speed
		car.battery -= car.batteryDrain

		return car
	}

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// panic("Please implement the CanFinish function")
	ref := car.battery / car.batteryDrain

	if ref * car.speed >= track.distance {
		return true
	} else {
		return false
	}
}
