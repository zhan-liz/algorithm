package leetcode

/*
104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return maxDepthD(root, 0)
}

func maxDepthD(root *TreeNode, num int) int {
	if root == nil {
		return num
	}
	num++
	lem := maxDepthD(root.Left, num)
	rigm := maxDepthD(root.Right, num)
	if lem >= rigm {
		return lem
	}
	return rigm
}
