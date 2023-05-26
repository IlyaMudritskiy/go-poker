package objects

import (
	"fmt"
	"github.com/ilyamudritskiy/go-poker/utils"
)

var cardBack = [9]string{}

func fillBack() {
	cardBack[0] = "┌─────────┐"
	for i := 1; i < 8; i++ {
		cardBack[i] = fmt.Sprintf("%s%s%s", "│", utils.Yellow.Get("░░░░░░░░░"), "│")
	}
	cardBack[8] = "└─────────┘"
}

// Card is a struct that represents a single card
//   - Suite: suite of the card
//   - Value: value of the card
//   - Face: []string is face of the card, filles automatically
//   - Back: []string is back of the card, filles automatically
type Card struct {
	Suite string
	Value string
	Face  []string
	Back  [9]string
}

// Print a card using mini format or ASCII
//   - "mini" - mini format (suite value)
//   - "big" - big format (11x9 ASCII card)
//   - "back" - back of the card (11x9 ASCII card)
func (card Card) Print(style string) {
	// Print mini version of cards
	if style == "mini" {
		fmt.Printf("%s %s  ", card.Suite, card.Value)
	}

	// Print ASCII version of face
	if style == "big" {
		for _, line := range card.Face {
			fmt.Println(line)
		}
	}

	// Print ASCII version of back
	if style == "back" {
		for _, line := range card.Back {
			fmt.Println(line)
		}
	}
}

// Creates the face of the card using suite and value
// For card with value 10 apllies different spacing
func getCardFace(suite string, value string) []string {
	var face = []string{}

	face = append(face, "┌─────────┐")

	if value == "10" {
		face = append(face, fmt.Sprintf("│ %s %s    │", value, suite))
	} else {
		face = append(face, fmt.Sprintf("│ %s  %s    │", value, suite))
	}

	face = append(face, "│         │")
	face = append(face, "│         │")
	face = append(face, fmt.Sprintf("│    %s    │", suite))
	face = append(face, "│         │")
	face = append(face, "│         │")

	if value == "10" {
		face = append(face, fmt.Sprintf("│    %s %s │", suite, value))
	} else {
		face = append(face, fmt.Sprintf("│    %s  %s │", suite, value))
	}

	face = append(face, "└─────────┘")

	return face
}

func color(suite string) string {
	if suite == Hearts {
		return utils.Red.Get(suite)
	}
	if suite == Diamonds {
		return utils.Red.Get(suite)
	}
	if suite == Clubs {
		return utils.Blue.Get(suite)
	}
	if suite == Spades {
		return utils.Blue.Get(suite)
	}
	return suite
}

// Function for getting ready Card with filled fields
func GetCard(suite string, value string) Card {
	fillBack()
	return Card{color(suite), value, getCardFace(color(suite), value), cardBack}
}
