package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
// productionRate: number of cars produced
// successRate: percent of success produced
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	percent := successRate / 100 // get percent

	return float64(productionRate) * percent
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	productionByMinute := float64(productionRate / 60)
	percent := successRate / 100 // get percent

	return int(productionByMinute * percent)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	total := 0
	const price = 10000
	const priceByGroup = 95000

	if carsCount >= 10 {
		groups := carsCount / 10
		carsCount -= groups * 10
		total += groups * priceByGroup
	}

	if carsCount > 0 {
		total += carsCount * price
	}

	return uint(total)
}
