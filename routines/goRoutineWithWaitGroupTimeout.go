package routines

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func GoRoutineWithWaitGroupTimeout() {
	log.Println("GoRoutineWithWaitGroupTimeout...")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		success := make(chan bool)
		go func() {
			res := add(3, 2)
			fmt.Println("Result: ", res)
			success <- true

		}()
		select {
		case <-success:
			break
		case <-time.After(time.Duration(3 * time.Second)):
			fmt.Println("Error timed out")
			break
		}
		wg.Done()
	}()
	wg.Wait()
}

//uncomment the below code to get timeout issue
func add(a, b int) int {
	//time.Sleep(time.Duration(4*time.Second))
	return a + b
}
