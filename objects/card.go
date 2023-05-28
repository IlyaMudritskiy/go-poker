package objects

import (
	"fmt"
	s "github.com/ilyamudritskiy/go-poker/settings"
)

// Card is a struct that represents a single card
//   - Suite: suite of the card
//   - Value: value of the card
//   - Face: []string is face of the card, filles automatically
//   - Back: []string is back of the card, filles automatically
//	 - Used: bool shows if card was used
//   - Open: bool shows if card is face up or down
type Card struct {
	Suite string
	Value string
	Face  []string
	Back  []string
	Used  bool
	Open  bool
}

var cardBack = []string{}

func fillBack() {
	cardBack = append(cardBack, "┌─────────┐")
	for i := 1; i < 8; i++ {
		cardBack = append(cardBack, fmt.Sprintf("%s%s%s", "│", s.BackColor.Get("░░░░░░░░░"), "│"))
	}
	cardBack = append(cardBack, "└─────────┘")
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

// Color the card suite based on Settings
func colorSuite(suite string) string {
	if suite == Hearts {
		return s.HeartsColor.Get(suite)
	}
	if suite == Diamonds {
		return s.DiamindsColor.Get(suite)
	}
	if suite == Clubs {
		return s.ClubsColor.Get(suite)
	}
	if suite == Spades {
		return s.SpadesColor.Get(suite)
	}
	return suite
}

// Function for getting ready Card with filled fields and colored suites
func GetCard(suite string, value string) Card {
	if len(cardBack) == 0 {
		fillBack()
	}
	return Card{
		Suite: colorSuite(suite),
		Value: value,
		Face:  getCardFace(colorSuite(suite), value),
		Back:  cardBack,
		Used:  false,
		Open:  false}
}
