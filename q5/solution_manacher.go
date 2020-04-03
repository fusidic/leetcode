package q5

import "math"

func longestPalindrome3(s string) string {
	origin := s
	s = pretreatment(s)
	len := len(s)

	// p[] 存储各个点上的中心扩展(回文)值
	p := make([]int, len)
	var center, bound, maxLen, centerIndex int

	p[0] = 0
	p[1] = 0
	center = 1

	// 对right点 确定p[right]
	for right := 1; right < len-1; right++ {
		left := 2*center - right

		// 超出边界
		if right > bound {
			p[right] = 0
		} else {
			// 取left，bound-i (到边界的距离)，p[left] (镜像位置)
			p[right] = int(math.Min(float64(bound-right), float64(p[left])))
		}

		// 对正在处理的right点，尝试中心扩展
		for s[right+1+p[right]] == s[right-1-p[right]] {
			p[right]++
		}

		// 超出界限, 更新 center 和 bound
		if (right + p[right]) > bound {
			center = right
			bound = right + p[right]
		}
	}

	for i := 1; i < len-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}
	startIndex := (centerIndex - maxLen) / 2
	return origin[startIndex : startIndex+maxLen]

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
