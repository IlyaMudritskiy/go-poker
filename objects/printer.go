package objects

import (
	"fmt"
	"strings"
)

//==============================================================================
//=                 Printing and converting to string for Card                 =
//==============================================================================

func (card Card) AsString(mini bool) string {
	var spacing string = " "

	if card.Value == "10" {
		spacing = ""
	}

	if mini {
		return fmt.Sprintf("%s%s%s\n", card.Suite, spacing, card.Value)
	} else {
		return faceAsString(card.getImage())
	}
}

func (card Card) Print(mini bool) {
	fmt.Print(card.AsString(mini))
}

//==============================================================================
//=                 Printing and converting to string for Hand                 =
//==============================================================================

func (hand Hand) AsString() string {
	var result string = ""

	if hand == nil {
		fmt.Println("Hand is empty")
		return result
	}

	result = cardsAsString(hand)

	return result
}

func (hand Hand) Print() {
	fmt.Print(hand.AsString())
}

//==============================================================================
//=                 Printing and converting to string for Deck                 =
//==============================================================================

func (deck Deck) AsString() string {
	var result string = ""

	result = result + "=================================| Deck |==================================\n"

	for i := 0; i < 4; i++ {
		var s string = ""

		for j := 0; j < 13; j++ {
			var spacing = "   "

			if deck.Cards[i*13+j].Value == "10" {
				spacing = "  "
			}

			s += fmt.Sprintf("%s %s%s", deck.Cards[i*13+j].Suite, deck.Cards[i*13+j].Value, spacing)
		}

		result = result + s + "\n"
	}

	result = result + "===========================================================================\n"

	return result
}

func (deck Deck) Print() {
	fmt.Print(deck.AsString())
}

//==============================================================================
//=                Printing and converting to string for Table                 =
//==============================================================================

func (table Table) AsString() string {
	return cardsAsString(table.Cards)
}

func (table Table) Print() {
	fmt.Print(table.AsString())
}

//==============================================================================
//=                            Supporting functions                            =
//==============================================================================

func (card Card) getImage() []string {
	if card.Open {
		return card.Face
	} else {
		return card.Back
	}
}

func faceAsString(lines []string) string {
	var result string = ""
	for _, line := range lines {
		result = result + strings.TrimSpace(line) + "\n"
	}
	return result
}

func cardsAsString(cards []Card) string {
	var result string = ""

	if len(cards) == 0 {
		fmt.Println("No cards were given")
		return result
	}

	for i := 0; i < 9; i++ {
		for _, card := range cards {
			result = result + strings.TrimSpace(card.getImage()[i]) + "  "
		}
		result = result + "\n"
	}

	return result
}
