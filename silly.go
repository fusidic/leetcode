package main

import (
	"fmt"
)

func main() {
	s := "babad"
	s = pretreatment(s)
	fmt.Println(s)
}

func pretreatment(s string) string {
	len := len(s)
	if len == 0 {
		return "^$"
	}
	cue := []rune("^#$")
	var res []rune
	runes := []rune(s)
	res = append(res, cue[0])
	for i := 0; i < len; i++ {
		res = append(res, cue[1], runes[i])
	}
	res = append(res, cue[1:]...)
	return string(res)
}
