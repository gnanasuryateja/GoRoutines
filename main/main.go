package main

import (
	"fmt"
	"log"

	rt "goroutines/routines"
)

func main() {
	log.Println("Go Routines...")
	rt.SimpleRoutine()
	rt.GoRoutineWithWaitGroup()
	//stopGoRoutineUsingChannel
	numbers := rt.Numbers()
	fmt.Println(<-numbers)
	fmt.Println(<-numbers)
	numbers <- 0
	rt.GoRoutineWithWaitGroupTimeout()
	rt.ForLoopUsingGoRouintes()
	rt.ForLoopUsingGoRouintesWithChannel()
}
