package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeetArray2Tree(t *testing.T) {
	tests := []struct {
		name      string
		leetArray []int
		want      []int
	}{
		{
			name:      "tree01",
			leetArray: []int{1, 2, NULL, NULL, 3},
			want:      []int{1, 2, NULL, NULL, 3},
		},
		{
			name:      "tree02",
			leetArray: []int{1, NULL, 2, 3, 4, 5, 6, 7, 8},
			want:      []int{1, NULL, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:      "tree03",
			leetArray: []int{1, NULL, 2, 3},
			want:      []int{1, NULL, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Tree2LeetArray(LeetArray2Tree(tt.leetArray)))
		})
	}
}
