package leetcode

import "math"

// Given a string, find the length of the longest substring without repeating characters.

// Example 1:

// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:

// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Example 3:

// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.

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
