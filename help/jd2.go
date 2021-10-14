package help

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

func testJd2() int {
	var a, b, f, k int
	fmt.Scanf("%d %d %d %d", &a, &b, &f, &k)
	return minTimes(a, b, f, k)
}
