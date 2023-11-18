// package birdwatcher
package main

import (
	"fmt"
)

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for i := 0; i < len(birdsPerDay); i++ {
		count += birdsPerDay[i]
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// perday[week-1*7:week] // start := if week == 1.... else (week - 1) * 7
	if week <= 0 {
		panic("Naughty, naughty: week has to be bigger than 0")
	}
	var startIndex int
	if week == 1 {
		startIndex = 0
	} else {
		startIndex = (week - 1) * 7
	}
	birdsInWeek := birdsPerDay[startIndex : startIndex + 7]
	var count int
	for i := 0; i < 7; i++ {
		count += birdsInWeek[i]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// newSlice := make([]int, 0, len(birdsPerDay)) // allocate memory and length with make !!!
	for i := range birdsPerDay {
		if i % 2 == 0 {
			birdsPerDay[i]++
		}
		// newSlice = append(newSlice, val)
	}
	// return newSlice
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(BirdsInWeek(birdsPerDay, 2)) // 12
	fmt.Println(BirdsInWeek(birdsPerDay, 1)) // 22
	// fmt.Println(BirdsInWeek(birdsPerDay, 0))
	fmt.Println(FixBirdCountLog(birdsPerDay))
}