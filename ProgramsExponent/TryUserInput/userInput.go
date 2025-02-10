package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Println("Whats your name?")
	fmt.Scan(&name)
	fmt.Printf("Your name: %v\n", name)

	myAge := 10
	newAge := makeMeOlder(myAge)
	fmt.Println(newAge)
	fmt.Println(myAge)

}

func makeMeOlder(age int) int {
	age += 5
	return age
}
