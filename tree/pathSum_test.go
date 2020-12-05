package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		inputTree []int
		sum       int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testcase01",
			args: args{
				inputTree: []int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, 5, 1},
				sum:       22,
			},
			want: [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}},
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, pathSum(structures.LeetArray2Tree(tt.args.inputTree), tt.args.sum))
	}
}
