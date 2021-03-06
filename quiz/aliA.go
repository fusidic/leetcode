package quiz

import "fmt"

// 三个线程，分别输出 A B C，并发输出 ABCABC

func multiPrint() {
	cnt := 5
	A := make(chan struct{})
	B := make(chan struct{})
	C := make(chan struct{})
	go func() {
		for true {
			s := <-C
			if cnt >= 0 {
				fmt.Println("A")
				cnt--
				A <- s
			} else {
				break
			}
		}
		return
	}()

	go func() {
		for true {
			s := <-A
			if cnt >= 0 {
				fmt.Println("B")
				cnt--
				B <- s
			} else {
				break
			}
		}
		return
	}()

	go func() {
		for true {
			s := <-B
			if cnt >= 0 {
				fmt.Println("C")
				cnt--
				C <- s
			} else {
				break
			}
		}
		return
	}()
	s := struct{}{}
	C <- s
}
