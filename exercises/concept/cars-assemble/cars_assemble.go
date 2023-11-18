// package cars
package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute, those are per hour
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount <= 0 {
		return 0
	}
	return uint((carsCount / 10)* 95000 + (carsCount % 10) * 10000)
}

func main() {
	fmt.Println(CalculateWorkingCarsPerMinute(426, 80))
	fmt.Print(CalculateCost(32)) // 
}

/*
func floatingPointEquals(got, want float64) bool {
	absoluteDifferenceBelowThreshold := math.Abs(got-want) <= floatEqualityThreshold
	relativeDifferenceBelowThreshold := math.Abs(got-want)/(math.Abs(got)+math.Abs(want)) <= floatEqualityThreshold
	return absoluteDifferenceBelowThreshold || relativeDifferenceBelowThreshold
}
The code snippet you provided is a method to compare two floating-point numbers (got and want) for equality
 within a certain threshold. Due to the nature of floating-point arithmetic, direct comparisons can be unreliable,
  so using thresholds to determine if two floats are "close enough" is a common technique.
*/
