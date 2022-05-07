package routines

import "log"

func GetMessage(ch chan string){
	log.Println("Inside go routine")
	ch<-"got the message"
	log.Println("Exiting go routine")
}
