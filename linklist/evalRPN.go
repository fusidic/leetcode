package linklist

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	var temp int
	for _, str := range tokens {
		switch str {
		case "+":
			temp = stack[len(stack)-1] + stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], temp)
		case "-":
			temp = stack[len(stack)-1] - stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], temp)
		case "*":
			temp = stack[len(stack)-1] * stack[len(stack)-2]
			stack = append(stack[:len(stack)-2], temp)
		case "/":
			temp = stack[len(stack)-2] / stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], temp)
		default:
			n, _ := strconv.Atoi(str)
			stack = append(stack, n)
		}
	}
	return stack[0]
}
