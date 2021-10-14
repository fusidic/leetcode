package string

func kmp(s1, s2 string) int {
	if len(s2) == 0 {
		return -1
	}
	next := getNext(s2)
	j := -1
	for i := 1; i < len(s1); i++ {
		for j >= 0 && s1[i] != s2[j+1] {
			j = next[j]
		}
		if s1[i] == s2[j+1] {
			j++
		}
		if j == len(s2)-1 {
			return i - len(s2) + 1
		}
	}
	return -1
}

func kmp2(s1, s2 string) int {
	if len(s2) == 0 {
		return -1
	}
	next := getNext(s2)
	j := -1
	for i := 1; i < len(s1); i++ {
		for j >= 0 && s1[i] != s2[j+1] {
			j = next[j]
		}
		if s1[i] == s2[j+1] {
			j++
		}
		if j == len(s2)-1 {
			return i - len(s2) + 1
		}
	}
	return -1
}

func getNext(s string) []int {
	n := len(s)
	next := make([]int, n)
	j := -1
	next[0] = j
	for i := 1; i < n; i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[j+1] == s[i] {
			j++
		}
		next[i] = j
	}
	return next
}

func getNext3(s string) []int {
	next := make([]int, len(s))
	j := -1
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j >= 0 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
