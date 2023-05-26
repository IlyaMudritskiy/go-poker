package objects

import (
	"fmt"
)

func PrintTable () {
	var s string
	var deskHand = Hand{GetCard(" ", " "), GetCard(" ", " "), GetCard(" ", " "), GetCard(" ", " "), GetCard(" ", " ")}
	for i, _ := range deskHand[0].Back {
		for _, card := range deskHand {
			s = s + card.Back[i] + "   "
		}
		s = s + "\n"
	}
	fmt.Println(s)
}