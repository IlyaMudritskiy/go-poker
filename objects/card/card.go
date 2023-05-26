package card

import "fmt"

type Card struct {
	Suite string
	Value string
	Face  []string
	Back  [9]string
}

var cardBack = [9]string{}

func GetCard(suite string, value string) Card {
	fillBack()
	return Card{suite, value, getCardFace(suite, value), cardBack}
}

func fillBack() {
	cardBack[0] = "┌─────────┐"
	for i := 1; i < 8; i++ {
		cardBack[i] = fmt.Sprintf("%s%s%s", "│", "░░░░░░░░░", "│")
	}
	cardBack[8] = "└─────────┘"
}

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
