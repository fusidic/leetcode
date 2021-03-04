package array

import (
	"fmt"
	"testing"
	"time"

	"gotest.tools/assert"
)

func Test_getMax(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{
			name: "num01",
			args: 8873,
			want: 8873,
		},
		{
			name: "num01",
			args: 1234,
			want: 4231,
		},
		{
			name: "num03",
			args: 1001,
			want: 1100,
		},
	}
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 2; i++ {
			ch <- 1
		}
		close(ch)
	}()
	time.Sleep(1000)
	go func() {
		for i := 0; i < 3; i++ {
			elem, _ := <-ch
			fmt.Println(elem)
			time.Sleep(1000)
		}
	}()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, GetMax(tt.args))
		})
	}
}
