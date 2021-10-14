package huawei

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type resSlice [][]int

func (s resSlice) Len() int      { return len(s) }
func (s resSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s resSlice) Less(i, j int) bool {
	var sum1, sum2 int

	for _, v := range s[i] {
		sum1 += v
	}

	for _, v := range s[j] {
		sum2 += v
	}

	return sum1 > sum2 || s[i][0] > s[j][0]
}

func inputHandlerHw2() {

	// 第一行输入
	var averageScore, n int
	fmt.Scanf("%d %d", &averageScore, &n)

	// 第二行输入
	sc := bufio.NewScanner(os.Stdin)
	nums := make([]int, n*2)
	sc.Scan()
	strs := strings.Split(sc.Text(), " ")
	for i, v := range strs {
		num, _ := strconv.Atoi(v)
		nums[i] = num
	}

	res := solve2(nums, averageScore)

	if res == nil {
		fmt.Println("0")
		return
	}

	resStr := []string{}

	for _, v := range res {
		resStr = append(resStr, strconv.Itoa(v))
	}

	fmt.Println(strings.Join(resStr, " "))
}

func solve2(nums []int, k int) []int {

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	// 构建 mod 二维数组，按余数分类
	mod := make([][]int, k)
	for i := 0; i < k; i++ {
		mod[i] = []int{}
	}
	for _, v := range nums {
		// 需要处理负数
		mod[(v%k+k)%k] = append(mod[(v%k+k)%k], v)
	}

	// 检验是否可以匹配成功
	for i := 1; i <= k/2; i++ {
		if len(mod[i]) != len(mod[k-i]) {
			return nil
		}
	}
	if k%2 == 0 {
		if len(mod[k/2])%2 != 0 {
			return nil
		}
	}

	var res resSlice

	if k%2 == 0 {
		for i := 1; i < k/2; i++ {
			for j := 0; j < len(mod[i]); j++ {
				res = append(res, orderPair(mod[i][j], mod[k-i][j]))
			}
		}

		for i := 0; i < len(mod[k/2]); i += 2 {
			res = append(res, orderPair(mod[k/2][i], mod[k/2][i+1]))
		}
	} else {
		for i := 1; i <= k/2; i++ {
			for j := 0; j < len(mod[i]); j++ {
				res = append(res, orderPair(mod[i][j], mod[k-i][j]))
			}
		}
	}

	for i := 0; i < len(mod[0]); i += 2 {
		res = append(res, orderPair(mod[0][i], mod[0][i+1]))
	}

	sort.Sort(res)
	trueRes := []int{}

	for _, v := range res {
		trueRes = append(trueRes, v...)
	}

	return trueRes
}

func orderPair(a, b int) []int {
	if a > b {
		return []int{a, b}
	}
	return []int{b, a}
}
