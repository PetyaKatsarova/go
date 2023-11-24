// package savings
package main

import "fmt"

// FixedInterestRate has a value of 5% (5/100)

const FixedInterestRate = 0.05

// GetFixedInterestRate returns the FixedInterestRate constant.
func GetFixedInterestRate() float32 {
	return float32(FixedInterestRate)
}

// DaysPerYear has a value of 365
const DaysPerYear = 365

// GetDaysPerYear returns the DaysPerYear constant.
func GetDaysPerYear() int {
	return DaysPerYear
}

// Jan-Dec have values of 1-12
const (
	Jan = iota + 1
	Feb
	Mar
	Apr
	May 
	Jun
	Jul
	Aug
	Sept
	Oct
	Nov
	Dec
)

// GetMonth returns the value for the given month.
func GetMonth(month int) int {
	return month
}

// AccNo type for a string - this is a stub type included to demonstrate how the untyped string constant can be used where this type is required.
type AccNo string

// AccountNo has a value of XF348IJ
const AccountNo = "XF348IJ"

// GetAccountNumber returns the AccountNo constant.
func GetAccountNumber() AccNo {
	return AccountNo
}


func main() {
	fmt.Println(GetDaysPerYear())
	fmt.Println(GetMonth(Dec))
	fmt.Println(GetAccountNumber())
}

/*

In Go, iota is a predeclared identifier used to create successive untyped integer constants, usually
 within a const block. It simplifies the definition of incrementing numbers. The value of iota starts at 
 0 within each const block and increments by one for each new constant.

 type Weekday int

// Use iota to define constants representing days of the week
const (
    Monday Weekday = iota // 0
    Tuesday               // 1
    Wednesday             // 2
    Thursday              // 3
    Friday                // 4
    Saturday              // 5
    Sunday                // 6
)

 const (
    Flag1 = 1 << iota // 1 (1 << 0)
    Flag2             // 2 (1 << 1)
    Flag3             // 4 (1 << 2)
    Flag4             // 8 (1 << 3)
)
*/
