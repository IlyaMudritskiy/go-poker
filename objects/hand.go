package objects

import "fmt"

type Hand []Card

func (hand Hand) Print(style string) {
	fmt.Println("Your hand:")

	if style == "mini" {
		for _, card := range hand {
			card.Print("mini")
		}
	}

	if style == "big" {
		for i := 0; i < len(hand[0].Face); i++ {
			fmt.Printf("%s   %s\n", hand[0].Face[i], hand[1].Face[i])
		}
	}

	if style == "back" {
		for i := 0; i < len(hand[0].Back); i++ {
			fmt.Printf("%s   %s\n", hand[0].Back[i], hand[1].Back[i])
		}
	}

	fmt.Println()
}
