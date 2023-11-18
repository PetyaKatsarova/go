// package cards
package main

import "fmt"

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) || index < 0 {
		slice = append(slice, value)
	} else {
		slice[index] = value
	}
	return slice
}
/*
-------- VARIADIC FUNCTIONS -------------------
*/

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index + 1:]...)
	// append in slice from 0 to i the slice from i+1 till end
}


func main() {
	fmt.Println(PrependItems( []int{5, 2, 10, 6, 8, 7, 0, 9}, 1))
	fmt.Println("removed at i 1: ", RemoveItem( []int{5, 2, 10, 6, 8, 7, 0, 9}, 1))
}