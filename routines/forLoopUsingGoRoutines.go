package routines

import (
	"fmt"
	"log"
	"sync"
)

func ForLoopUsingGoRouintes() {
	log.Println("ForLoopUsingGoRouintes...")
	slice := []string{"a", "b", "c", "d", "e"}
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running for loop…")
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			val := slice[i]
			fmt.Printf("i: %v, val: %v\n", i, val)
		}(i)
	}
	wg.Wait()
	fmt.Println("Finished ForLoopUsingGoRouintes")
}

func ForLoopUsingGoRouintesWithChannel() {
	log.Println("ForLoopUsingGoRouintesWithChannel...")
	slice := []string{"a", "b", "c", "d", "e"}
	sliceLength := len(slice)
	var wg sync.WaitGroup
	wg.Add(sliceLength)
	fmt.Println("Running for loop…")
	ch := make(chan string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		go func(i int) {
			defer wg.Done()
			ch <- slice[i]
		}(i)
	}
	wg.Wait()
	close(ch)
	for val := range ch {
		fmt.Println(val)
	}
	fmt.Println("Finished ForLoopUsingGoRouintesWithChannel")
}
