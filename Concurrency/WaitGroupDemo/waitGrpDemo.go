package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("Wait groups demo...")
	fmt.Println("Total go routines after main:", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go f1(&wg)
	//time.Sleep(time.Millisecond)
	go f2(&wg)
	fmt.Println("Total go routines with two functions:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Total go routines after finishing two functions execution:", runtime.NumGoroutine())
	fmt.Println("Execution complete")
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 function started")
	for i := 0; i < 5; i++ {
		fmt.Println("i = ", i)
	}
	time.Sleep(time.Second)
	fmt.Println("f1 function completed")
	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("f2 function started")
	for i := 5; i < 10; i++ {
		fmt.Println("i = ", i)
	}
	time.Sleep(time.Second)
	fmt.Println("f2 function completed")
	wg.Done()
}
