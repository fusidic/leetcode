package pdf

import (
	"log"
	"strings"
	"sync"
)

func printNumAnndLetter() {
	letter, number := make(chan bool), make(chan bool)
	wait := sync.WaitGroup{}
	stop := make(chan struct{})
	go func() {
		i := 1
		for {
			select {
			case <-number:
				log.Print(i)
				i++
				log.Print(i)
				i++
				letter <- true
			case <-stop:
				log.Print("num fin")
				return
			}

		}
	}()
	wait.Add(1)
	go func(wait *sync.WaitGroup) {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					return
				}
				log.Print(str[i : i+1])
				i++
				if i >= strings.Count(str, "")-1 {
					wait.Done()
					stop <- struct{}{}
					return
				}
				log.Print(str[i : i+1])
				i++
				number <- true
			}

		}
	}(&wait)

	number <- true
	wait.Wait()
}
