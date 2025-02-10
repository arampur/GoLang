package main

import (
	"fmt"
)

type Person struct {
	fName string
	lName string
}

func (person Person) fullName() string {
	return person.fName + " " + person.lName
}

// structs as pointer
func (person *Person) changeName(newFName string) {
	person.fName = newFName
}

func main() {
	p1 := Person{fName: "Amith", lName: "Ramp"}
	fmt.Println("Before changing first name")
	fmt.Println(p1.fullName())
	fmt.Println("After changing first name")
	p1.changeName("Akshath")
	fmt.Println(p1.fullName())

	var myArray [3]int
	fmt.Println(myArray)

}
