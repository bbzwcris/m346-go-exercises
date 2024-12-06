package main

import "fmt"

var firstName, lastName string = "Cristian", "Antinori"

// firstName = "Cristian"
// lastName = "Antinori"
var dayOfBirth, monthOfBirth, yearOfBirth int = 12, 7, 2007

// dayOfBirth = 12
// monthOfBirth = 7
// yearOfBirth = 2007
var numberOfSiblings int = 2
var heightInMeters float32 = 1.69
var zodiacSign string = "♋︎"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %s\n", zodiacSign)
}
