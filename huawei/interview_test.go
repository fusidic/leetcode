package huawei

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	testcases := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "case0",
			matrix: [][]int{{0}},
			want:   [][]int{{0}},
		},
		{
			name:   "case1",
			matrix: [][]int{{1, 2}, {4, 3}},
			want:   [][]int{{4, 1}, {3, 2}},
		},
		{
			name:   "case2",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "case3",
			matrix: [][]int{{1, 0, 0, 2}, {0, 0, 0, 0}, {0, 0, 0, 0}, {4, 0, 0, 3}},
			want:   [][]int{{4, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {3, 0, 0, 2}},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, rotate(tt.matrix))
		})
	}
}
