package main
// This code is not dry and not rational but for the practise of interfaces

import (
	"fmt"
)

func main() {
		// Create a slice of Greeter interface types
		greeters := []Greeter {
			Bulgarian{},
			Italian{},
			// PortugeseGreeter{},
			Italian{},
		}
	
		// Names for each language
		names := []string{"Petka", "Luiggi", "Tomaso Giulio Micheli"}
	
		// Iterate over each Greeter and greet in their language
		for i, greeter := range greeters {
			msg := SayHello(names[i], greeter)
			fmt.Println(msg)
		}
}

type Bulgarian struct{}
func (b Bulgarian) LanguageName() string { return "Bulgarian" }
func (b Bulgarian) Greet(name string) string { return "Zdravei, " + name + "!" }

type Italian struct{}
func(i Italian) LanguageName() string {return "Italian" }
func(i Italian) Greet(name string) string { return "Ciao " +name + "!"}

// type PortugeseGreeter struct{}
// func(i PortugeseGreeter) LanguageName() string {return "Portugese" }
// func(i PortugeseGreeter) Greet(name string) string { return "Ola " +name + "!"}
