package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
input:
	a,1,5
	b,3,5
	c,1,3

output:
	b c
*/

// 可用区
const available = (1 << 4) - 1

// 乘车区间售价
func getCosts() [5]int {
	return [5]int{0, 100, 180, 240, 280}
}

type Client struct {
	name string
	from int
	to   int
	cost int
	// 用位表示乘车区间
	mask int
}

func ticketSystem() string {

	// 读取输入
	sc := bufio.NewScanner(os.Stdin)
	clients := []Client{}
	for sc.Scan() {
		if len(sc.Text()) == 0 {
			break
		}
		clients = append(clients, getClient(sc.Text()))
	}
	res := solve1(clients)
	return strings.Join(res, " ")
}

func getClient(info string) Client {
	strs := strings.Split(info, ",")
	from, _ := strconv.Atoi(strs[1])
	to, _ := strconv.Atoi(strs[2])
	return Client{
		name: strs[0],
		from: from,
		to:   to,
		cost: getCosts()[to-from],
		mask: (1 << (to - 1)) - (1 << (from - 1)),
	}
}

// 0-1 背包解法
func solve1(clients []Client) []string {
	booked := make([][]bool, len(clients)+1)
	for i := 0; i < len(clients)+1; i++ {
		booked[i] = make([]bool, available+1)
	}

	// 记录最大价值
	dp := make([][]int, len(clients)+1)
	for i := 0; i < len(clients)+1; i++ {
		dp[i] = make([]int, available+1)
	}

	for i := 1; i <= len(clients); i++ {
		client := clients[i-1]
		for j := 0; j <= available; j++ {
			fmt.Printf("mask: %d cal: %d\n", client.mask, client.mask&j)

			if (client.mask & j) != client.mask {
				// 不符合乘坐区间
				dp[i][j] = dp[i-1][j]
				continue
			}
			if dp[i-1][j] <= (dp[i][j & ^client.mask] + client.cost) {

				dp[i][j] = dp[i][j & ^client.mask] + client.cost
				booked[i][j] = true
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp)

	// 获取最优组合
	bits := available
	res := []string{}
	for i := len(clients); i > 0; i-- {
		client := clients[i-1]

		if booked[i][bits] {
			res = append([]string{client.name}, res...)
			bits &= ^client.mask
		}
	}

	return res
}
