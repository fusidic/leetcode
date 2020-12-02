package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_isSubStructure(t *testing.T) {
	type args struct {
		treeA []int
		treeB []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase01",
			args: args{
				treeA: []int{1, 2, 3, 4},
				treeB: []int{3},
			},
			want: true,
		},
		{
			name: "testcase02",
			args: args{
				treeA: []int{1, 2, 3},
				treeB: []int{3, 1},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isSubStructure(structures.LeetArray2Tree(tt.args.treeA), structures.LeetArray2Tree(tt.args.treeB)))
		})
	}
}
