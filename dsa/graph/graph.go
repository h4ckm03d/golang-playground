package graph

import "math"

type Node struct {
	Value string
	Edges []*Edge
}

type Edge struct {
	Start *Node
	End   *Node
	Cost  int
}

type Graph struct {
	Nodes []*Node
}

func (g *Graph) AddNode(value string) *Node {
	node := &Node{
		Value: value,
	}

	g.Nodes = append(g.Nodes, node)

	return node
}

func (g *Graph) AddEdge(start, end *Node, cost int) {
	edge := &Edge{
		Start: start,
		End:   end,
		Cost:  cost,
	}

	start.Edges = append(start.Edges, edge)
}

func (n *Node) Neighbors() []*Node {
	neighbors := []*Node{}

	for _, edge := range n.Edges {
		neighbors = append(neighbors, edge.End)
	}

	return neighbors
}

func (g *Graph) getCost(start, end *Node) int {
	for _, edge := range start.Edges {
		if edge.End == end {
			return edge.Cost
		}
	}
	return -1
}
func (g *Graph) MinCost(start, end *Node) int {
	visited := make(map[*Node]bool)
	costs := make(map[*Node]int)

	// initialize all costs to infinity except for the starting node
	for _, node := range g.Nodes {
		costs[node] = math.MaxInt32
	}
	costs[start] = 0

	for {
		// find the unvisited node with the smallest cost
		var node *Node
		minCost := math.MaxInt32
		for _, n := range g.Nodes {
			if !visited[n] && costs[n] < minCost {
				minCost = costs[n]
				node = n
			}
		}

		if node == nil {
			break // no path found
		}

		visited[node] = true
		for _, neighbor := range node.Neighbors() {
			if !visited[neighbor] {
				cost := costs[node] + g.getCost(node, neighbor)
				if cost < costs[neighbor] {
					costs[neighbor] = cost
				}
			}
		}
	}

	// return the cost of the path from start to end
	return costs[end]
}
