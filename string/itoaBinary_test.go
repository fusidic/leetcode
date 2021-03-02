package string

import (
	"fmt"
	"testing"
)

func Test_ConvertToBin(t *testing.T) {
	fmt.Println(convertToBin(10, 8))
	fmt.Println(convertToBin(-10, 8))
}
