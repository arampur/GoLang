package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Println("Whats your name?")
	fmt.Scan(&name)
	fmt.Printf("Your name: %v\n", name)

	fmt.Println("Whats your age?")
	fmt.Scan(&age)
	newAge := makeMeOlder(age)
	fmt.Println("Your current age:", age)
	fmt.Println("After 5 years you will be:", newAge)

}

func makeMeOlder(age int) int {
	age += 5
	return age
}
