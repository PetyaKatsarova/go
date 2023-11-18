// package chance
package main

import (
	"crypto/rand"
	"math/big"
	"fmt"
)

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	n, err := rand.Int(rand.Reader, big.NewInt(20))
	if err != nil { panic(err) }
	return int(n.Int64()) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	const precision = 1000000 
	n, err := rand.Int(rand.Reader, big.NewInt(precision)) // from 0 to 1 milion - 1
	if err != nil {
		return 0.0 // might be better to add return err, too
	}
	return (float64(n.Int64()) / precision) * 12.0
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}

	// can use rand.Shuffle but it is more fun to use Fisher_Yates shuffle, called Knuth too
	for i := len(animals) - 1; i > 0; i-- {
		randIndexBig, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		randIndex := int(randIndexBig.Int64())
		animals[i], animals[randIndex] = animals[randIndex], animals[i]
	}

	return animals
}


// TEST 1
// func main() {
// 	fmt.Println("wtf")
// 	for i := 0; i < 10; i++ {
// 		n := RollADie()
// 		fmt.Printf("[%d] the num: %d\n", i, n) 
// 	}
// }

//TEST 2
func main() {
	for i := 0; i < 10; i++ {
		n := GenerateWandEnergy()
		fmt.Printf("[%d] the num: %f\n", i, n) // between 0.0 and 12.0
	}
}