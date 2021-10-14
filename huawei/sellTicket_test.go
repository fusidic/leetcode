package huawei

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve1(t *testing.T) {

	type args struct {
		inputs []string
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case1",
			args: args{
				inputs: []string{"a,1,5", "b,1,3", "c,3,5"},
			},
			want: "b c",
		},
		{
			name: "case2",
			args: args{
				inputs: []string{"a,1,4", "b,2,3", "c,1,3", "d,3,4"},
			},
			want: "c d",
		},
		{
			name: "case3",
			args: args{
				inputs: []string{"a,2,3", "b,1,3", "c,3,4", "d,4,5", "e,1,3"},
			},
			want: "c d e",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			clients := []Client{}
			for _, v := range tt.args.inputs {
				clients = append(clients, getClient(v))
			}
			assert.Equal(t, tt.want, strings.Join(solve1(clients), " "))
		})
	}
}
