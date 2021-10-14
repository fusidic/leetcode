package huawei

import (
	"testing"

	"gotest.tools/assert"
)

func Test_solve2(t *testing.T) {

	type args struct {
		average int
		n       int
		nums    []int
	}

	testCases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{
				average: 5,
				n:       5,
				nums:    []int{1, 10, 5, 4, 3, 2, 7, 6, 8, -1},
			},
			want: []int{10, 5, 8, 7, 6, 4, 3, 2, 1, -1},
		},
		{
			name: "case2",
			args: args{
				average: 1,
				n:       3,
				nums:    []int{2, 4, 6, 8, 10, 12},
			},
			want: []int{12, 10, 8, 6, 4, 2},
		},
		{
			name: "case3",
			args: args{
				average: 10,
				n:       3,
				nums:    []int{2, 3, 4, 7, 8, 9},
			},
			want: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.DeepEqual(t, tt.want, solve2(tt.args.nums, tt.args.average))
		})
	}

}
