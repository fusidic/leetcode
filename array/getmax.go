package array

import (
	"strconv"
)

// GetMax 给定一个正整数，可最多交换一次此正整数中的两位数字
func GetMax(x int) int {
	str := []byte(strconv.Itoa(x))
	// 存放 0～9 的索引
	idx := make([]int, 10)
	// 记录每个元素的索引值，并更新到最右边
	for i, c := range str {
		idx[c-'0'] = i
	}
	// 左起遍历
	for i, c := range str {
		// 找到存在的最大位的位置
		for j := 9; j > int(c-'0'); j-- {
			if idx[j] > i {
				// 交换
				str[i], str[idx[j]] = str[idx[j]], str[i]
				res, _ := strconv.Atoi(string(str))
				return res
			}
		}
	}
	return x
}

func main() {
	arr1 := 12345
	// arr2 := 23424
	// arr3 := 8873
	GetMax(arr1)
}

// 测试
// func Test_getMax(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		args int
// 		want int
// 	}{
// 		{
// 			name: "num01",
// 			args: 8873,
// 			want: 8873,
// 		},
// 		{
// 			name: "num01",
// 			args: 1234,
// 			want: 4231,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			assert.Equal(t, tt.want, GetMax(tt.args))
// 		})
// 	}
// }
