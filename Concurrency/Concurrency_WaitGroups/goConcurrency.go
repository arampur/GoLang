package main

//When you execute the below program, you will see that the f1 is not executed meaning This is a simple function f1 will never be called.
//Reason for this: every standalone go application will create only one go routine which is main go routine.
//Here without waitgroup main function doesnt wait for the go routine to execute/complete. Main terminates before even it gets to the go routine which is why f1 is never executed.

import (
	"fmt"
	"runtime"
	"sync"
	"time"
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

//The below is the output for this program..

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

//Solving above problem using waitgroups

func goroutineDemo() {
	var wg sync.WaitGroup
	wg.Add(2) //Number of go routines to wait
	//go f1() //Uncommenting this will give you output as above.
	go wgFunc(&wg)

	wg.Wait()
	fmt.Println("Execution is done")
}

func wgFunc(wg *sync.WaitGroup) {
	fmt.Println("F1 function started...")
	time.Sleep(time.Second)
	wg.Done() //(*wg).Done() - both mean the same
}

/*
Go Routines Demo..
f1 total go routines: 2
f2 function
i = 3
i = 4
This is a simple function f1
i = 5
i = 6
f2 is complete
*/

//The above output f1 function is not consistent and you will notice when you run the program multiple times, the results are different each time.
//This output depends on which goroutines finish first. This scenario of unsynchornized access to the shared memory is called Data Race.
//Data Race occurs when two go routines access memory at the same time.
