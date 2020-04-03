package leetutils

// Reverse the string and return
func Reverse(s string) string {
	runes := []rune(s)
	for head, tail := 0, len(runes)-1; head < tail; head, tail := head+1, tail-1 {
		runes[head], runes[tail] = runes[tail], runes[head]
	}
	return string(runes)
}
