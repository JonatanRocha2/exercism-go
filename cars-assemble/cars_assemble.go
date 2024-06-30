package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(productionRate)) * (float64(successRate) / 100)
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCatsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCatsPerHour / 60)
}

func CalculateCost(carsCount int) uint {
	return (uint(carsCount / 10 * 95000)) + (uint(carsCount % 10 * 10000))
}
