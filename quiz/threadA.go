package quiz

import (
	"log"
)

func printer(order int, printQ chan int, stopC chan<- bool) {
	for {
		value := <-printQ
		if value == 50 {
			stopC <- true
			return
		}
		value++
		log.Printf("Thread %d print: %d", order, value)
		printQ <- value
	}
}

func autoPrint() {
	printQ := make(chan int)
	stopC := make(chan bool)
	go printer(1, printQ, stopC)
	go printer(2, printQ, stopC)
	printQ <- 0
	log.Println("stop", <-stopC)
}
