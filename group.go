package lilraft

import "fmt"

// Group ...
type Group struct {
	nodes map[string]*node
}

// Run ...
func (g *Group) Run() {
	g.nodes = map[string]*node{
		"abc123": &node{
			id: "abc123",
		},
	}

	for id := range g.nodes {
		fmt.Printf("State of %s: %d\n", g.nodes[id].id, g.nodes[id].state)
	}
}

func (g *Group) send() {}