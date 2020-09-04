package main

import (
	"fmt"
	"strconv"
)

// globally/package scoped variable, need to use full declaration syntax (no :=)
var name string = "Davey"

// Animal struct, may be exported
type Animal struct {
	Name   string // `required max:"100"`
	Origin string
}

// Bird struct with Animal characteristics, may be exported
type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	// function/block scoped variables
	var (
		msg         string
		age         int
		daveySmells bool = true
		ageString   string
		piString    string
		pungency    int
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

	// fmt.Printf("pi: %v, %T\n", pi, pi)
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

	fmt.Printf("fine: %v\n", fine)
	fmt.Printf("malodorous: %v\n", malodorous)
	fmt.Printf("acrid: %v\n", acrid)
	fmt.Printf("putrid: %v\n", putrid)
	fmt.Printf("rancid: %v\n", rancid)

	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	fmt.Println(statePopulations)
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia")

	if popGA, isGeorgia := statePopulations["Georgia"]; isGeorgia {
		fmt.Printf("Georgia pop.: %v\n", popGA)
	}
	if popOH, isOhio := statePopulations["Ohio"]; isOhio {
		fmt.Printf("Ohio pop.: %v\n", popOH)
	}

	emu := Bird{
		Animal:   Animal{Name: "Emu", Origin: "Australia"},
		SpeedKPH: 48,
		CanFly:   false,
	}

	fmt.Println(emu)
	// tag := reflect.TypeOf(Animal{})
	// field, _ := tag.FieldByName("Name")
	// fmt.Println(field.Tag)
}
