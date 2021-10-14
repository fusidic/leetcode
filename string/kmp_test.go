package string

import (
	"testing"

	"gotest.tools/assert"
)

func TestKmp(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}

	testcases := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				s1: "aabaabaaf",
				s2: "aabaaf",
			},
			want: 3,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, testcase.want, kmp(testcase.args.s1, testcase.args.s2))
		})
	}
}
