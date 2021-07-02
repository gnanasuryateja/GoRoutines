package routines

import (
	"fmt"
	"log"
)

func SimpleRoutine() {
	log.Println("SimpleRoutine...")
	go func() {
		fmt.Println("Executing SimpleRoutine")
	}()
	fmt.Scanln()
}
