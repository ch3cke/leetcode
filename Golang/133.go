package main

func cloneGraph(node *Node) *Node {
	visit := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if node == nil {
			return node
		}
		if _, ok := visit[node]; ok {
			return visit[node]
		}
		cloneNode := &Node{node.Val, []*Node{}}
		visit[node] = cloneNode
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
		}
		return cloneNode
	}
	return cg(node)
}
