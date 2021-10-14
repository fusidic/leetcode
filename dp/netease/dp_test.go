package netease

import (
	"fmt"
	"log"
	"testing"
)

func Test_minSailCost(t *testing.T) {

	input := [][]int{{1, 0, 1, 1, 0}, {2, 2, 0, 1, 0}, {1, 1, 2, 1, 1}, {0, 2, 0, 0, 1}}
	fmt.Println(minSailCost(input))
	log.Println(minSailCost(input))
}
