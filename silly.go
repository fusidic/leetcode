package main

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	head := 0
	rear := 0
	longest := 0
	dict := make(map[string]int)
	for index, v := range s {
		head = index
		_, ok := dict[string(v)]
		if ok {
			rear = int(math.Max(float64(dict[string(v)]+1), float64(rear)))
			// for ch, num := range dict {
			// 	if num < rear {
			// 		delete(dict, ch)
			// 	}
			// }
		}
		dict[string(v)] = index
		// longest = int(math.Max(float64(longest), float64(head-rear+1)))
		longest = int(math.Max(float64(longest), float64(head-rear+1)))
	}
	return longest
}

func main() {
	s := "tmmzuxt"
	r := lengthOfLongestSubstring(s)
	fmt.Println(r)
}
