package quiz

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main1() {
	var set int

	fmt.Scan(&set)
	ans := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < set; i++ {
		scanner.Scan()
		bound := numbers(scanner.Text())
		scanner.Scan()
		nums := numbers(scanner.Text())
		// fmt.Println(getNum(nums, bound[1]))
		ans = append(ans, getNum(nums[0:bound[0]], bound[1]))
	}
	for _, v := range ans {
		fmt.Println(v)
	}
}

func numbers(s string) []int {
	var n []int
	for _, f := range strings.Fields(s) {
		i, _ := strconv.Atoi(f)
		n = append(n, i)
	}
	return n
}

func getNum(nums []int, k int) int {
	cnt := 0
	start := nums[0]
	nums = append(nums, -1<<32)
	i := 0
	for i < len(nums) {
		if cnt == k {
			break
		}
		if nums[i] == start {
			start++
			i++
		} else {
			start++
			cnt++
		}
	}
	return start - 1
}
