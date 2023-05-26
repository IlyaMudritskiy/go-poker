package main

import (
	"github.com/ilyamudritskiy/go-poker/objects"
)

func main() {
	var deck = objects.GetDeck()
	objects.PrintTable()
	var hand = deck.DealHand(0)
	hand.Print("big")
	
}
