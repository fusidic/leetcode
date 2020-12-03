package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_levelOrder2(t *testing.T) {
	tests := []struct {
		name      string
		inputTree []int
		want      [][]int
	}{
		{
			name:      "tree01",
			inputTree: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
			want:      [][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, levelOrder2(structures.LeetArray2Tree((tt.inputTree))))
	}
}
