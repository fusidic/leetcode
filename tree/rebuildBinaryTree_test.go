package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preOrder []int
		inOrder  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "buildTree",
			args: args{
				preOrder: []int{1, 2, 3},
				inOrder:  []int{2, 3, 1},
			},
			want: []int{1, 2, structures.NULL, structures.NULL, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, structures.Tree2LeetArray(buildTree(tt.args.preOrder, tt.args.inOrder)))
		})
	}
}
