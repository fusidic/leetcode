package huawei

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindRadius(t *testing.T) {
	testcases := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			name:  "case1",
			nums1: []int{1, 2, 3},
			nums2: []int{2},
			want:  1,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findRadius(tt.nums1, tt.nums2))
		})
	}
}
