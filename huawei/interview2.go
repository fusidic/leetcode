package huawei

type Node struct {
	value int
	flag  bool
	next  []*Node
}

func travel(root *Node) []int {

	res := []int{}

	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil || node.flag {
			return
		}

		res = append(res, node.value)
		node.flag = true
		for _, n := range node.next {
			dfs(n)
		}

	}

	dfs(root)
	return res
}
