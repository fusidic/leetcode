package quiz

import (
	"testing"
)

func Test_flowmedian(t *testing.T) {
	operations := [][]int{{1, 5}, {2}, {1, 3}, {2}, {1, 6}, {2}, {1, 7}, {2}}
	flowmedian(operations)
}
