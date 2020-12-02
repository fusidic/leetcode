package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		leetOrder []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "tree01",
			args: args{
				leetOrder: []int{3, 9, 20, structures.NULL, structures.NULL, 15, 7},
			},
			want: []int{3, 9, 20, 15, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := structures.LeetArray2Tree(tt.args.leetOrder)
			res := levelOrder(root)
			assert.Equal(t, tt.want, res)
		})
	}
}
