package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		name    string
		inputT  []int
		outputL [][]int
	}{
		{
			name:   "testcase01",
			inputT: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
			outputL: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structures.LeetArray2Tree(tt.inputT)
			assert.Equal(t, tt.outputL, zigzagLevelOrder(root))
		})
	}
}
