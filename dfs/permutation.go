package dfs

// Permutation ...
func Permutation(s string) []string {
	str := []byte(s)
	res := []string{}
	length := len(str)
	var dfs func(x int)
	dfs = func(x int) {
		if x == length-1 {
			res = append(res, string(str))
			return
		}
		dict := map[byte]bool{}
		for i := x; i < length; i++ {
			if _, ok := dict[str[i]]; ok {
				continue
			}
			str[x], str[i] = str[i], str[x]
			dfs(x + 1)
			str[i], str[x] = str[x], str[i]
		}
	}
	dfs(0)
	return res
}
