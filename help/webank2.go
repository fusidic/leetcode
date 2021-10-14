package help

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mainHelp() {
	var n, m, x, y int
	fmt.Scanf("%d %d %d %d", n, m, x, y)

	matrix := [][]int{}

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		strs := strings.Split(sc.Text(), " ")
		nums := make([]int, m)
		for i, v := range strs {
			num, _ := strconv.Atoi(v)
			nums[i] = num
		}
		matrix = append(matrix, nums)
	}

	sc.Scan()
	move := strings.Split(sc.Text(), " ")
	x1, y1 := 0, 0
	res := 0
	for i := 0; i < len(move); i++ {
		c := move[i]
		if c == "l" {
			x1++
			if i == 0 {
				if x1 >= m || matrix[y1][x1] == -1 {
					res += y
				} else {
					res += max(matrix[y1][x1], matrix[y1][x1-1])
				}
			} else {
				if move[i-1] != c {
					if x1 >= m || matrix[y1][x1] == -1 {
						res += (x + y)
					} else {
						res = res + x + max(matrix[y1][x1], matrix[y1][x1-1])
					}
				} else {
					if x1 >= m || matrix[y1][x1] == -1 {
						res += y
					} else {
						res = res + max(matrix[y1][x1], matrix[y1][x1-1])
					}
				}
			}
		}

		if c == "h" {
			x1 -= 1
			if i == 0 {
				if x1 < 0 || matrix[y1][x1] == -1 {
					res += y
				} else {
					res = res + max(matrix[y1][x1], matrix[y1][x1+1])
				}
			} else {
				if move[i-1] != c {
					if x1 < 0 || matrix[y1][x1] == -1 {
						res = res + x + y
					} else {
						res = res + x + max(matrix[y1][x1], matrix[y1][x1+1])
					}
				} else {
					if x1 < 0 || matrix[y1][x1] == -1 {
						res += y
					} else {
						res = res + max(matrix[y1][x1], matrix[y1][x1+1])
					}
				}
			}
		}

		if c == "k" {
			y1 -= 1
			if i == 0 {
				if y1 < 0 || matrix[y1][x1] == -1 {
					res += y
				} else {
					res = res + max(matrix[y1][x1], matrix[y1+1][x1])
				}
			} else {
				if move[i-1] != c {
					if y1 < 0 || matrix[y1][x1] == -1 {
						res = res + x + y
					} else {
						res = res + x + max(matrix[y1][x1], matrix[y1+1][x1])
					}
				} else {
					if y1 < 0 || matrix[y1][x1] == -1 {
						res += y
					} else {
						res = res + max(matrix[y1][x1], matrix[y1+1][x1])
					}
				}
			}
		}

		if c == "j" {
			y1 += 1
			if i == 0 {
				if y1 >= n || matrix[y1][x1] == -1 {
					res = res + y
				} else {
					res = res + max(matrix[y1][x1], matrix[y1-1][x1])
				}
			} else {
				if move[i-1] != c {
					if y1 >= n || matrix[y1][x1] == -1 {
						res = res + x + y
					} else {
						res = res + x + max(matrix[y1][x1], matrix[y1-1][x1])
					}
				} else {
					if y1 >= n || matrix[y1][x1] == -1 {
						res = res + y
					} else {
						res = res + max(matrix[y1][x1], matrix[y1-1][x1])
					}
				}
			}
		}
	}

	fmt.Println(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
