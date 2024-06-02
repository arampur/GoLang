// PROGRAM TO RESOLVE DATA RACE CONDITION...

package main

import (
	"fmt"
	"sync"
	"time"
)

//In the below program any code between the lock and unlock will be executed by only one go routine at a time.
//Can also be written as
//m.Lock()
//n++
//m.Unlock()
//wg.Done()

func main() {

	const gr = 100
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(gr * 2) //200 go routines

	var n int = 0
	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n++
			wg.Done()
		}() //function invoke, not doing this will result in expression in go must be function call error.

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n--
			wg.Done()
		}()
	}
	fmt.Println("value of n: ", n)
}

//Executing the above program will now give value of n as 0 all the time.
