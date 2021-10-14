// package main

// import (
// 	"log"
// 	"strings"
// 	"sync"
// )

// func main() {
// 	printNumAnndLetter()
// }

// func printNumAnndLetter() {
// 	letter, number := make(chan bool), make(chan bool)
// 	wait := sync.WaitGroup{}
// 	stop := make(chan struct{})
// 	go func() {
// 		i := 1
// 		for {
// 			select {
// 			case <-number:
// 				log.Print(i)
// 				i++
// 				log.Print(i)
// 				i++
// 				letter <- true
// 			case <-stop:
// 				log.Print("num fin")
// 				return
// 				// default:
// 				// 	fmt.Print("default1")
// 				// 	break
// 			}
// 		}
// 	}()
// 	wait.Add(1)
// 	go func(wait *sync.WaitGroup) {
// 		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 		i := 0
// 		for {
// 			<-letter
// 			if i >= strings.Count(str, "")-1 {
// 				wait.Done()
// 				return
// 			}
// 			log.Print(str[i : i+1])
// 			i++
// 			if i >= strings.Count(str, "")-1 {
// 				wait.Done()
// 				stop <- struct{}{}
// 				return
// 			}
// 			log.Print(str[i : i+1])
// 			i++
// 			number <- true
// 		}
// 	}(&wait)

// 	number <- true
// 	wait.Wait()

// }

package main

import "fmt"

func minTimes(a, b, f, k int) int {

	if b < (f-a)*2 || b < f*2 {
		return -1
	}

	times := 0

	// 第一次
	rest := b - f
	if rest < 2*(a-f) {
		rest = b
		times++
	}

	for i := 2; i < k; i++ {
		if i%2 == 1 {
			rest -= 2 * f
			if rest < 2*(a-f) {
				rest = b
				times++
			}
		} else {
			rest -= 2 * (a - f)
			if rest < 2*f {
				rest = b
				times++
			}
		}
	}

	// 最后一次做单独判断
	if k%2 == 1 {
		rest -= 2 * f
		if rest < (a - f) {
			rest = b
			times++
		}
	} else {
		rest -= 2 * (a - f)
		if rest < f {
			rest = b
			times++
		}
	}

	return times
}

func findDog(likes, prices, costs []int, n, k, d int) int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if likes[j] > likes[j+1] {
				likes[j], likes[j+1] = likes[j+1], likes[j]
				prices[j], prices[j+1] = prices[j+1], prices[j]
				costs[j], costs[j+1] = costs[j+1], costs[j]
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if prices[i] > k {
			continue
		}
		if float64(d*(i+1))/float64(n) < float64(costs[i]) {
			continue
		}
		return likes[i]
	}

	return -1
}

func main() {
	// var a, b, f, k int
	// fmt.Scanf("%d %d %d %d", &a, &b, &f, &k)
	// fmt.Println(minTimes(a, b, f, k))
	n, k, d := 5, 100, 100
	likes := []int{1, 2, 3, 4, 5}
	prices := []int{30, 60, 90, 120, 150}
	costs := []int{11, 22, 44, 88, 176}

	fmt.Println(findDog(likes, prices, costs, n, k, d))

}
