package cars

const costPerIndividualCar = 10_000
const costPerGroupOfTenCars = 95_000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsOfTenCars := carsCount / 10
	remainingCarsOutOfGroups := carsCount % 10

	return uint(groupsOfTenCars*costPerGroupOfTenCars + remainingCarsOutOfGroups*costPerIndividualCar)
}
