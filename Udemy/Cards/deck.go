package main

import "fmt"

// Create a new type of deck, which is a slice of strings
type deck []string

func (d deck) print() { // use a one or two letter for the var
	for i, card := range d { // to represent the type of the receiver
		fmt.Println(i, card)
	}
}
