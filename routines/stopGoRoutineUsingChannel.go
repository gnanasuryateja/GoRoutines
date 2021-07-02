package routines

import "log"

func Numbers() chan int {
	log.Println("StopGoRoutineUsingChannel...")
	ch := make(chan int)
	go func() {
		n := 1
		for {
			select {
			case ch <- n:
				n++
			case <-ch:
				return
			}
		}
	}()
	return ch
}
