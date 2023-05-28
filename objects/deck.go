package objects

import (
	"fmt"
	"math/rand"
	"time"

	s "github.com/ilyamudritskiy/go-poker/settings"
)

type Deck struct {
	Cards   []Card
	Current int // Points to the first card available in the deck
}

// Creating a deck of cards (52)
func GetDeck() Deck {
	var deck Deck

	for _, suite := range suites {
		for _, value := range cardValues {
			deck.Cards = append(deck.Cards, GetCard(suite, value))
		}
	}

	return deck.randomise()
}

// Mixes the deck in the reandom order
func (deck Deck) randomise() Deck {
	// Set new seed
	rand.Seed(time.Now().UnixNano())

	// Get a slice of random permutations
	rand_order := rand.Perm(len(deck.Cards))

	var result = Deck{}
	result.Current = 0

	// Fill result slice according to order in rand_order
	for _, i := range rand_order {
		result.Cards = append(result.Cards, deck.Cards[i])
	}

	return result
}

// Get 1 card from the Deck noting that card is used
func (deck *Deck) Draw() Card {
	if deck.Current > len(deck.Cards) {
		fmt.Println("No more cards available!")
	}

	var card = deck.Cards[deck.Current]

	// Set value color to UsedColor to check when deck is printed
	deck.Cards[deck.Current].Value = s.UsedColor.Get(deck.Cards[deck.Current].Value)

	// Temporarily here
	card.Open = true
	// Note that card is used
	card.Used = true

	deck.Current++

	return card
}
