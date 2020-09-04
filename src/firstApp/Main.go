package main

import (
	"fmt"
	"strconv"
)

var name string = "Davey"

func main() {
	var (
		msg         string
		age         int
		daveySmells bool = true
	)

	const (
		// _ = iota // _ = iota or errorValue = iota if a zero value is not wanted
		fine = iota
		malodorous
		acrid
		putrid
		rancid
		pi float32 = 3.141592654
	)

	var (
		ageString string
		piString  string
		pungency  int
	)

	// fmt.Printf("%v, %T\n", pi, pi)
	msg = fmt.Sprintf("%v smells, I mean, rocks!", name)
	age = 48
	pungency = rancid
	ageString = strconv.Itoa(age)
	piString = strconv.FormatFloat(float64(pi), 'g', -1, 32) // see https://golang.org/pkg/strconv/#example_FormatFloat

	fmt.Println("Checking pungency...")
	if daveySmells && pungency >= acrid {
		fmt.Println(msg)
		fmt.Printf("%v has smelled, I mean, rocked for %v years!\n", name, ageString)
		fmt.Printf("And %v really loves apple and cherry %v\n", name, piString)
	} else {
		fmt.Println("Pungency level tolerable!")
	}

	// fmt.Printf("%v\n", fine)
	// fmt.Printf("%v\n", malodorous)
	// fmt.Printf("%v\n", acrid)
	// fmt.Printf("%v\n", putrid)
	// fmt.Printf("%v\n", rancid)
}
