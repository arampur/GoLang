package main

//When you execute the below program, you will see that the f1 is not executed meaning This is a simple function f1 will never be called.
//Reason for this: every standalone go application will create only one go routine which is main go routine.
//Here main function doesnt wait for the go routine to execute/complete. Main terminates before even it gets to the go routine which is why f1 is never executed.

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Go Routines Demo..")
	//To create a go routine
	go f1() //This command creates a goroutine
	fmt.Println("f1 total go routines:", runtime.NumGoroutine())

	f2() //This is a normal function call. Not a goroutine
}

func f1() {
	fmt.Println("This is a simple function f1")
}

func f2() {
	fmt.Println("f2 function")
	for i := 3; i < 7; i++ {
		fmt.Println("i =", i)
	}
	fmt.Println("f2 is complete")
}

//The below is the output for this file..

/*
Go Routines Demo..
f1 go routines 2
f2 function
i = 3
i = 4
i = 5
i = 6
f2 is complete
*/
