package main

import (
	o "github.com/ilyamudritskiy/go-poker/objects"
)

func main() {
	var deck = o.GetDeck()
	deck.Print()

	deck.Cards[0].Print(true)
	
	var hand = o.Hand{deck.Draw(), deck.Draw()}
	hand.Print()

	var table = o.GetTable()
	table.AddCard(deck.Draw())
	table.AddCard(deck.Draw())
	table.AddCard(deck.Draw())
	table.Print()
	deck.Print()
}
