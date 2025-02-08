package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	return this.deserializeHelper(&nodes)
}

func (this *Codec) deserializeHelper(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	if (*nodes)[0] == "null" {
		*nodes = (*nodes)[1:]
		return nil
	}
	val, _ := strconv.Atoi((*nodes)[0])
	root := &TreeNode{Val: val}
	*nodes = (*nodes)[1:]
	root.Left = this.deserializeHelper(nodes)
	root.Right = this.deserializeHelper(nodes)
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */