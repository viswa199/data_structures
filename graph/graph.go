package graph

import "fmt"

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key      string
	Adjacent []*Vertex
}

func (g *Graph) AddVertex(key string) error {
	if !Contains(g.Vertices, key) {
		g.Vertices = append(g.Vertices, &Vertex{Key: key})
		return nil
	}
	return fmt.Errorf("%v ", "DUPLICATE VALUES ARE NOT ALLOWED")
}

func (g *Graph) Print() {
	for _, v := range g.Vertices {
		fmt.Printf("\nVertex %v:", v.Key)
		for _, a := range v.Adjacent {
			fmt.Printf("%v ", a.Key)
		}
	}
}

func Contains(s []*Vertex, key string) bool {
	for _, v := range s {
		if v.Key == key {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to string) error {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("NODES ARE NOT FOUND.")
		return err
	} else if Contains(fromVertex.Adjacent, to) {
		err := fmt.Errorf("EDGE ALREADY IS THERE.")
		return err
	} else {
		fromVertex.Adjacent = append(fromVertex.Adjacent, toVertex)
		return nil
	}
}

func (g *Graph) getVertex(key string) *Vertex {
	for _, v := range g.Vertices {
		if v.Key == key {
			return v
		}
	}
	return nil
}

func (g *Graph) NodeCount() int {
	return len(g.Vertices)
}