package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_recoverTree(t *testing.T) {
	tests := []struct {
		name    string
		inputT  []int
		outputT []int
	}{
		{
			name:    "testcase01",
			inputT:  []int{1, 3, structures.NULL, structures.NULL, 2},
			outputT: []int{3, 1, structures.NULL, structures.NULL, 2},
		},
		{
			name:    "testcase02",
			inputT:  []int{3, 1, 4, structures.NULL, structures.NULL, 2},
			outputT: []int{2, 1, 4, structures.NULL, structures.NULL, 3},
		},
		{
			name:    "testcase03",
			inputT:  []int{2, 3, 1},
			outputT: []int{2, 1, 3},
		},
		{
			name:    "testcase04",
			inputT:  []int{146, 71, -13, 55, structures.NULL, 231, 399, 321, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, -33},
			outputT: []int{146, 71, 321, 55, structures.NULL, 231, 399, -13, structures.NULL, structures.NULL, structures.NULL, structures.NULL, structures.NULL, -33},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structures.LeetArray2Tree(tt.inputT)
			recoverTree(root)
			assert.Equal(t, tt.outputT, structures.Tree2LeetArray(root))
		})

	}
}
