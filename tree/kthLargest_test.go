package tree

import (
	"testing"

	"github.com/fusidic/leetcode/structures"
	"github.com/stretchr/testify/assert"
)

func Test_kthLargest(t *testing.T) {
	type args struct {
		inputBST []int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase01",
			args: args{
				inputBST: []int{3, 1, 4, structures.NULL, 2},
				k:        1,
			},
			want: 4,
		},
		{
			name: "testcase02",
			args: args{
				inputBST: []int{5, 3, 6, 2, 4, structures.NULL, structures.NULL, 1},
				k:        3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, kthLargest(structures.LeetArray2Tree(tt.args.inputBST), tt.args.k))
	}
}
