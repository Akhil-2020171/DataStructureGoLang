package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func clone(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if _, ok := visited[node]; ok {
		return visited[node]
	}
	newNode := &Node{Val: node.Val}
	visited[node] = newNode
	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(neighbor, visited))
	}
	return newNode
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	return clone(node, visited)
}