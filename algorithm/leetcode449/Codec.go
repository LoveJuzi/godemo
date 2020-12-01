package leetcode449

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// TreeNode 结点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec BST定义
type Codec struct {
}

// Constructor 构建BST
func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {

	var cnt int

	T := buildT()

	T = append(T, root)

	var rt string

	for len(T) != 0 {
		rt += itoa(T[0].Val)
		rt += ","
		if T[0].Left != nil {
			cnt++
			rt += itoa(cnt)
			T = append(T, T[0].Left)
		} else {
			rt += "0"
		}
		rt += ","
		if T[0].Right != nil {
			cnt++
			rt += itoa(cnt)
			T = append(T, T[0].Right)
		} else {
			rt += "0"
		}
		rt += ";"
		T = T[1:]
	}

	return rt
}

func buildT() []*TreeNode {
	return make([]*TreeNode, 0)
}

type Node struct {
	v1 int
	v2 int
	v3 int
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	var v1, v2, v3 int
	T := make([]Node, 0)
	for len(data) != 0 {
		data, v1, v2, v3 = f1(data)
		T = append(T, Node{v1, v2, v3})
	}

	return f2(T)
}

func f1(s string) (string, int, int, int) {
	var i1, i2, i3 int
	for i1 = 0; s[i1] != ','; i1++ {
	}

	for i2 = i1 + 1; s[i2] != ','; i2++ {
	}

	for i3 = i2 + 1; s[i3] != ';'; i3++ {
	}

	return s[i3+1:], atoi(s[0:i1]), atoi(s[i1+1 : i2]), atoi(s[i2+1 : i3])
}

func f2(T []Node) *TreeNode {
	if len(T) == 0 {
		return nil
	}

	T2 := make([]*TreeNode, 0)
	for i := 0; i < len(T); i++ {
		T2 = append(T2, &TreeNode{T[i].v1, nil, nil})
	}
	for i := 0; i < len(T); i++ {
		if T[i].v2 != 0 {
			T2[i].Left = T2[T[i].v2]
		}
		if T[i].v3 != 0 {
			T2[i].Right = T2[T[i].v3]
		}
	}
	return T2[0]
}

func itoa(n int) string {
	return strconv.Itoa(n)
}

func atoi(s string) int {
	rt, _ := strconv.Atoi(s)
	return rt
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
