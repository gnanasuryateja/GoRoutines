package routines

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func printHello() {
	fmt.Println("Hello")
}

func printWorld() {
	fmt.Println("World")
}

func GoRoutineWithWaitGroup() {
	log.Println("GoRoutineWithWaitGroup...")
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		printHello()
		wg.Done()
	}()
	time.Sleep(time.Duration(1 * time.Millisecond))
	go func() {
		printWorld()
		wg.Done()
	}()
	wg.Wait()
}
