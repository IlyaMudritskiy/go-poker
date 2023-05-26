package objects

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck []Card

// Creating a deck of cards (52)
func GetDeck() Deck {
	var deck Deck

	for _, suite := range suites {
		for _, value := range cardValues {
			deck = append(deck, GetCard(suite, value))
		}
	}

	return deck.randomise()
}

// Mixes the deck in the reandom order
func (deck Deck) randomise() Deck {
	// Set new seed
	rand.Seed(time.Now().UnixNano())

	// Get a slice of random permutations
	rand_order := rand.Perm(len(deck))

	var result = []Card{}

	// Fill result slice according to order in rand_order
	for _, i := range rand_order {
		result = append(result, deck[i])
	}

	return result
}

// Returns a Hand (slice with two cards) and fills dealed cards in deck with whitespace (" ")
func (deck Deck) DealHand(start int) Hand {
	// Check if we try to deal cards out of deck length
	if start > len(deck)-2 {
		fmt.Println("Less than 2 cards are left. Please, get new deck.")
		return Hand{}
	}

	// Check if we dealed cards at index start
	if deck[start].Suite == " " || deck[start].Value == " " {
		fmt.Println("Current card is empty!")

		// Check if we dealed cards at index start + 1
		if deck[start+1].Suite == " " || deck[start+1].Value == " " {
			fmt.Println("Next card is empty!")
			return Hand{}
		}
		fmt.Println()
		
		return Hand{}
	}

	// Create hand to deal
	var hand = Hand{deck[start], deck[start+1]}

	// Fill dealed cards with whitespace (" ")
	for i := 0; i < start+1; i++ {
		deck[i].Suite = " "
		deck[i].Value = " "
		deck[i+1].Suite = " "
		deck[i+1].Value = " "
	}

	return hand
}

// Prints out the whole deck
func (deck Deck) Print() {
	fmt.Printf("=================================| Deck |==================================\n")

	for i := 0; i < 4; i++ {
		var s = ""

		for j := 0; j < 13; j++ {
			var spacing = "   "

			if deck[i*13+j].Value == "10" {
				spacing = "  "
			}

			s += fmt.Sprintf("%s %s%s", deck[i*13+j].Suite, deck[i*13+j].Value, spacing)
		}

		fmt.Println(s)
	}

	fmt.Printf("===========================================================================\n\n")
}
