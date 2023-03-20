package graph

import "testing"

func TestGraph_GetCost(t *testing.T) {
	g := &Graph{}

	// create nodes
	a := g.AddNode("A")
	b := g.AddNode("B")
	c := g.AddNode("C")

	// add edges with costs
	g.AddEdge(a, b, 2)
	g.AddEdge(b, c, 3)

	// test getCost
	cost := g.getCost(a, b)
	if cost != 2 {
		t.Errorf("got %d, expected %d", cost, 2)
	}

	cost = g.getCost(b, c)
	if cost != 3 {
		t.Errorf("got %d, expected %d", cost, 3)
	}
}

func TestGraph_DFS(t *testing.T) {
	g := &Graph{}

	// create nodes
	a := g.AddNode("A")
	b := g.AddNode("B")
	c := g.AddNode("C")
	d := g.AddNode("D")

	// add edges with costs
	g.AddEdge(a, b, 1)
	g.AddEdge(a, c, 2)
	g.AddEdge(b, d, 3)
	g.AddEdge(c, d, 4)

	// test MinCost
	cost := g.MinCost(a, d)
	if cost != 4 {
		t.Errorf("got %d, expected %d", cost, 4)
	}
}
