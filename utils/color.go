package utils

import (
	"fmt"
)

type Color string

var Reset Color = "\033[0m"
var Red Color = "\033[31m"
var Green Color = "\033[32m"
var Yellow Color = "\033[33m"
var Blue Color = "\033[34m"
var Purple Color = "\033[35m"
var Cyan Color = "\033[36m"
var White Color = "\033[37m"

func (color Color) Print(text string) {
	fmt.Println(color, text, Reset)
}

func (color Color) Get(text string) string {
	return fmt.Sprint(color, text, Reset)
}
