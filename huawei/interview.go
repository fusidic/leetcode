package huawei

func rotate(m [][]int) [][]int {
	l := len(m)
	for i := 0; i <= l/2; i++ {
		if l%2 == 1 && l/2 == i {
			continue
		}

		// 旋转逻辑
		for j := i; j < l-i-1; j++ {
			rightBound := l - i - 1

			tmp1 := m[i+j][rightBound]

			// 右上 m[i+j][rightBound]
			m[i+j][rightBound] = m[i][i+j]

			// 右下 m[rightBound][rightBound-j]
			tmp1, m[rightBound][rightBound-j] = m[rightBound][rightBound-j], tmp1

			// 左下 m[rightBound-j][i]
			tmp1, m[rightBound-j][i] = m[rightBound-j][i], tmp1

			// 左上 m[i][i+j]
			m[i][i+j] = tmp1

		}

	}

	return m
}
