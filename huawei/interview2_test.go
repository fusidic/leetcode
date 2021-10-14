package huawei

import (
	"fmt"
	"testing"
)

func TestTravel(t *testing.T) {
	testcases := []struct {
		name  string
		input Node
		want  []int
	}{
		{
			name:  "case1",
			input: Node{},
			want:  []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range testcases {
		fmt.Println(tt)
	}
}
