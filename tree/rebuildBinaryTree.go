package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

// method with stack
func buildTree(preorder []int, inorder []int) *structures.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	stack := list.New()
	root := &structures.TreeNode{}
	root.Val = preorder[0]
	cur := root
	j := 0
	for i := 1; i < len(preorder); i++ {
		// 判断 preorder 中相邻两元素的关系
		if cur.Val != inorder[j] {
			// 与 inorder 中不一致(非父节点)，相邻后点为左子树
			cur.Left = &structures.TreeNode{}
			cur.Left.Val = preorder[i]
			// 将该左子树存入栈中，便于之后确定其右子树
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			// 判断该点为父节点，开始回归
			j++
			for stack.Len() != 0 && stack.Back().Value.(*structures.TreeNode).Val == inorder[j] {
				cur = stack.Back().Value.(*structures.TreeNode)
				stack.Remove(stack.Back())
				j++
			}
			cur.Right = &structures.TreeNode{}
			cur.Right.Val = preorder[i]
			cur = cur.Right
		}
	}
	return root
}

// Java Code

// public TreeNode buildTree(int[] preorder, int[] inorder) {
//     if (preorder.length == 0) {
//         return null;
//     }
//     Stack<TreeNode> roots = new Stack<TreeNode>();
//     int pre = 0;
//     int in = 0;
//     //先序遍历第一个值作为根节点
//     TreeNode curRoot = new TreeNode(preorder[pre]);
//     TreeNode root = curRoot;
//     roots.push(curRoot);
//     pre++;
//     //遍历前序遍历的数组
//     while (pre < preorder.length) {
//         //出现了当前节点的值和中序遍历数组的值相等，寻找是谁的右子树
//         if (curRoot.val == inorder[in]) {
//             //每次进行出栈，实现倒着遍历
//             while (!roots.isEmpty() && roots.peek().val == inorder[in]) {
//                 curRoot = roots.peek();
//                 roots.pop();
//                 in++;
//             }
//             //设为当前的右孩子
//             curRoot.right = new TreeNode(preorder[pre]);
//             //更新 curRoot
//             curRoot = curRoot.right;
//             roots.push(curRoot);
//             pre++;
//         } else {
//             //否则的话就一直作为左子树
//             curRoot.left = new TreeNode(preorder[pre]);
//             curRoot = curRoot.left;
//             roots.push(curRoot);
//             pre++;
//         }
//     }
//     return root;
// }
