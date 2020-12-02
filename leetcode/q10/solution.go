package leetcode

import (
	"fmt"
	"log"
	"strings"
)

func debug(v ...interface{}) {
	log.Println("mermaid:")
	log.Println(v...)
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%v", i)
	case string:
		return fmt.Sprintf("%v", i)
	case bool:
		return fmt.Sprintf("%v", i)
	default:
		return fmt.Sprintf("%p", i)
	}
}

func isMatch(s string, p string) bool {
	begin := new(Node)
	// 头节点'>'，代表开始，不算在p中
	begin.C = '>'
	begin.Size = generatePattern(begin, p, 0)
	debug(begin.String())
	return check(begin, s, 0)
}

// Node ...
type Node struct {
	C        byte
	Parent   *Node
	Children map[byte][]*Node
	End      bool // 是否为终点
	Size     int  // 最小长度
}

func (n *Node) String() string {
	return n.StringLevel(0, make(map[*Node]bool))
}

// StringLevel ...
func (n *Node) StringLevel(level int, finishNodes map[*Node]bool) string {
	r := make([]string, 0)
	if n.End {
		r = append(r, fmt.Sprintf("  id%s{%v};", toString(n), string(n.C)))
	} else {
		r = append(r, fmt.Sprintf("  id%s(%v);", toString(n), string(n.C)))
	}
	finishNodes[n] = true
	for k, v := range n.Children {
		for _, c := range v {
			if _, ok := finishNodes[c]; !ok {
				r = append(r, c.StringLevel(level+1, finishNodes))
			}
			r = append(r, fmt.Sprintf("  id%s -- %s --> id%s;", toString(n), string(k), toString(c)))
		}
	}
	return strings.Join(r, "\n")
}

// Append ...
func (n *Node) Append(c byte, child *Node) {
	// m为map[byte][]*Node
	m := n.Children

	// 首节点Childre为nil，为其初始化
	if m == nil {
		m = make(map[byte][]*Node)
		n.Children = m
	}

	// list保存对应c的[]*Node
	list := m[c]
	if list == nil {
		// list为空，新建一个[]*Node
		list = make([]*Node, 0)
	}

	// 将child添加到Children中，已经有了就不加了
	for _, v := range list {
		if v == child {
			// child已经在list里面，不太懂为什么要搞出个list来
			m[c] = list
			return
		}
	}
	// child还没有添加到Children [c][]*Node中，添加进去
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *Node, str string, idx int) int {
	// 退出
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	// now 保存当前节点，vnow 指向下一节点
	vnow := now

	// 对下一个字符断言
	switch str[idx] {
	case '*':
		// *：将当前字节填充，匹配最小长队为0 (s*)
		now.Size = 0
		now.Append(now.C, now)
	default:
		// 普通字符，对下个字符创建Node (s)
		node := new(Node)
		node.C = str[idx]
		// 下一node，接到Children
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	// 对新更新的 Child 继续找孩子
	ret := generatePattern(vnow, str, idx+1)

	// 以'*'结尾，对上层返回0
	if ret == 0 {
		now.End = true
	}
	// ！！！
	addParent := now
	for addParent.Parent != nil {
		if addParent.Size == 0 {
			debug(toString(vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		} else {
			break
		}
	}
	return now.Size + ret
}

// check str[idx] 存在于 now.Children中，则返回true
func check(now *Node, str string, idx int) bool {
	// 匹配到句末，证明之前所有的字符都匹配上了
	if len(str) <= idx {
		return now.End
	}
	// list用来匹配字符串str，
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
		list = append(list, v)
	}
	for _, v := range list {
		// 在liat所有*Node的children中匹配下一个字符
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}
