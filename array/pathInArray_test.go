package array

import (
	"testing"

	"gotest.tools/assert"
)

func Test_Exist(t *testing.T) {
	// [["C","A","A"],["A","A","A"],["B","C","D"]]
	// "AAB"
	type args struct {
		matrix [][]byte
		word   string
	}
	testcase := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test01",
			args: args{
				matrix: [][]byte{
					{'C', 'A', 'A'},
					{'A', 'A', 'A'},
					{'B', 'C', 'D'},
				},
				word: "AAB",
			},
			want: true,
		},
	}

	for _, tt := range testcase {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Exist(tt.args.matrix, tt.args.word))
		})
	}
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		assert.Equal(t, tt.want, isSubStructure(structures.LeetArray2Tree(tt.args.treeA), structures.LeetArray2Tree(tt.args.treeB)))
	// 	})
	// }

}
