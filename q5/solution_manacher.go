package q5

func longestPalindrome3(s string) string {

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
	res = append(res, cue[1], cue[2])
	return string(res)
}
