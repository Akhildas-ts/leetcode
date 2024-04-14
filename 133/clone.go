package main

type Node struct {
    Val       int
    Neighbors []*Node
}

var visited map[*Node]*Node

func cloneGraph(node *Node) *Node {
    visited = make(map[*Node]*Node)
    return clone(node)
}

func clone(node *Node) *Node {
    if node == nil {
        return nil
    }

    if _, ok := visited[node]; ok {
        return visited[node]
    }

    cloneNode := &Node{Val: node.Val, Neighbors: make([]*Node, 0)}
    visited[node] = cloneNode

    for _, neighbor := range node.Neighbors {
        cloneNode.Neighbors = append(cloneNode.Neighbors, clone(neighbor))
    }

    return cloneNode
}

func main() {
    // Example usage
    node1 := &Node{Val: 1}
    node2 := &Node{Val: 2}
    node3 := &Node{Val: 3}

    node1.Neighbors = []*Node{node2, node3}
    node2.Neighbors = []*Node{node1, node3}
    node3.Neighbors = []*Node{node1, node2}

    clonedGraph := cloneGraph(node1)

    // Now you can use clonedGraph, which is a deep copy of the original graph
}
