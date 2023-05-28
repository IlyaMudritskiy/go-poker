package objects

import (
	"fmt"
)

type Table struct {
	Cards   []Card
	Current int
}

func GetTable() Table {
	var table = Table{}
	table.Current = 0

	for i := 0; i < 5; i++ {
		table.Cards = append(table.Cards, GetCard(" ", " "))
	}
	return table
}

func (table *Table) AddCard(card Card) {
	if table.Current == 5{
		fmt.Println("Can't add more cards - max cards already on the table")
		return
	} else {
		table.Cards[table.Current] = card
		table.Current++
	}
}
