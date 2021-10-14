package help

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputJd1() {
	var n, k, d int
	fmt.Scanf("%d %d %d", &n, &k, &d)
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	likes := toNums(sc.Text())
	sc.Scan()
	prices := toNums(sc.Text())
	sc.Scan()
	costs := toNums(sc.Text())
	fmt.Println(findDog(likes, prices, costs, n, k, d))
}

func toNums(s string) []int {
	strs := strings.Split(s, " ")
	res := []int{}
	for _, b := range strs {
		value, _ := strconv.Atoi(b)
		res = append(res, value)
	}

	return res
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

	res := -1
	for i := n - 1; i >= 0; i-- {
		if prices[i] > k {
			continue
		}
		if float64(d*(i+1))/float64(n) < float64(costs[i]) {
			continue
		}
		res = likes[i]
	}

	return res
}
