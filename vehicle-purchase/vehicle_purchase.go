package purchase

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	// panic("NeedsLicense not implemented")
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	// panic("ChooseVehicle not implemented")

	msg := "is clearly the better choice."

	if option1 <= option2 {
		return fmt.Sprintf("%v %s", option1, msg)
	} else {
		return fmt.Sprintf("%v %s", option2, msg)
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
// assume if the vehicle is less than 3 years old, it costs 80% of the original price
// If it is at least 10 years old, it costs 50%.
// If the vehicle is at least 3 years old but less than 10 years, it costs 70% of the original price.
func CalculateResellPrice(originalPrice, age float64) float64 {
	// panic("CalculateResellPrice not implemented")

	if int(age) < 3 {
		return originalPrice * 0.8
	} else if int(age) >= 10 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}
