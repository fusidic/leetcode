package huawei

import (
	"fmt"
	"sort"

	"github.com/fusidic/leetcode/common"
)

/**
input:
workstations: 1 3 5
sterilizers: 2
**/

func findRadius(workstations, sterilizers []int) int {
	sort.Ints(workstations)
	sort.Ints(sterilizers)
	sterilizers = append(sterilizers, 1<<32)
	sterilizers = append([]int{-1 << 32}, sterilizers...)

	var res int
	for _, w := range workstations {
		l, r := 0, len(sterilizers)

		for l < r {
			mid := (l + r) / 2
			if sterilizers[mid] < w {
				l = mid + 1
			} else {
				r = mid
			}
		}

		res = common.Max(res, common.Min(sterilizers[l]-w, w-sterilizers[l-1]))
	}
	return res
}

func inputHandlerHw1() {
	workstations := []int{1, 3, 5}
	sterilizers := []int{2}
	fmt.Println(findRadius(workstations, sterilizers))
}
