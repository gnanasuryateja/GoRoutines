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
	//communicateGoRoutineUsingChannel
	ch:=make(chan string)
	go rt.GetMessage(ch)
	fmt.Println(<-ch)
	rt.GoRoutineWithWaitGroupTimeout()
	rt.ForLoopUsingGoRouintes()
	rt.ForLoopUsingGoRouintesWithChannel()
}
