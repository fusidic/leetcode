package quiz

import "fmt"

func farthest(nums []int) {

}

func main() {
	var n int
	fmt.Scan(&n)
	set(n)
}

func set(n int) {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println(a)
}

// 这个是可行的
