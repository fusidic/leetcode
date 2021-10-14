package netease

import (
	"fmt"
	"log"
	"testing"
)

func Test_papers(t *testing.T) {

	input := []int{4, 4, 5}
	fmt.Println(papers(input))
	log.Println(papers(input))
}
