// package blackjack

package main

import (
	"fmt"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
		case "ace": return 11
		case "two": return 2
		case "three": return 3
		case "four": return 4
		case "five": return 5
		case "six": return 6
		case "seven": return 7
		case "eight": return 8
		case "nine": return 9
		case "ten", "jack", "queen", "king": return 10
		default: return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sumMyCards := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case card1 == "ace" && card2 == "ace": return "P"
	case sumMyCards == 21: 
		if dealerScore < 10 {
			return "W"
		} else {
			return "S" 
		}
	case sumMyCards >= 17 && sumMyCards <= 20: return "S"
	case sumMyCards >= 12 && sumMyCards <= 17 && dealerScore >= 7: return "H"
	case sumMyCards >= 12 && sumMyCards <= 18 && dealerScore < 7: return "S"
	case sumMyCards <= 11: return "H"
	default: return "H"
	}
}

func main() {
	fmt.Println(ParseCard("ace"))
	fmt.Println(ParseCard("bla"))
	fmt.Println(ParseCard(""))
	fmt.Println(ParseCard("three"))
	fmt.Println("expected P:", FirstTurn("ace", "ace", "")) // P
	fmt.Println("expected S:", FirstTurn("two", "ten", "")) // S
}