package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2) //200 go routines

	var n int = 0
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			n++
			wg.Done()
		}() //function invoke, not doing this will result in expression in go must be function call error.

		go func() {
			time.Sleep(time.Second / 10)
			n--
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("value of n: ", n)
}

//If you run this program, the output of n will be different in each run. which is due to condition called data race.
/*
go run dataRace.go
value of n:  3

go run dataRace.go
value of n:  -1

go run dataRace.go
value of n:  -3
*/

// If you run go run -race filename.go, you will see the below data race message on the console.
// This is one way to find out if your program caused data race condition or not. We will fix this condition through mutex.
/*
==================
WARNING: DATA RACE
Read at 0x00c000110038 by goroutine 105:
  main.main.func2()
  //path to your file

Previous write at 0x00c000110038 by goroutine 93:
  main.main.func2()
  //path to your file

Goroutine 105 (running) created at:
  main.main()
		//path to your file

Goroutine 93 (finished) created at:
  main.main()
      //path to your file
==================
value of n:  -1
Found 1 data race(s)
exit status 66
*/
