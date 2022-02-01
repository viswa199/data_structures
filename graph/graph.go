package graph

import (
	"fmt"
	"math"
)

type Graph struct {
	Vertices []*Vertex
}

type Vertex struct {
	Key           interface{}
	Adjacent      map[*Vertex]uint16
	IncomingNodes int
}

func (graph *Graph) AddVertex(key interface{}) {
	graph.Vertices = append(graph.Vertices, &Vertex{Key: key})
}

func (graph *Graph) Print() {
	for i, v := range graph.Vertices {
		fmt.Printf("%d:%v\n", i+1, v.Key)
	}
}

func (graph *Graph) Neighbours(key interface{}) {
	node := graph.contains(key)
	if node == nil {
		panic("data doesn't exist")
	}
	for keys := range node.Adjacent {
		fmt.Println(keys.Key)
	}
}

func (graph *Graph) AddEdge(source, destination interface{}, weight uint16) error {
	source_node := graph.contains(source)
	destination_node := graph.contains(destination)
	destination_node.IncomingNodes += 1

	if source_node != nil && destination_node != nil {
		if source_node.Adjacent == nil {
			source_node.Adjacent = make(map[*Vertex]uint16)
		}
		for keys := range source_node.Adjacent {
			if keys == destination_node {
				return fmt.Errorf("%v edge is already there.\n", destination)
			}
		}
		source_node.Adjacent[destination_node] = weight
	}
	return nil
}

func (graph Graph) contains(key interface{}) *Vertex {
	for i := 0; i < len(graph.Vertices); i++ {
		if graph.Vertices[i].Key == key {
			return graph.Vertices[i]
		}
	}
	return nil
}

func (graph *Graph) createVisited() map[*Vertex]bool {
	visited := make(map[*Vertex]bool)
	for i := 0; i < len(graph.Vertices); i++ {
		visited[graph.Vertices[i]] = false
	}
	return visited
}

func (graph *Graph) createUnVisited() map[*Vertex]bool {
	unvisited := make(map[*Vertex]bool)
	for i := 0; i < len(graph.Vertices); i++ {
		unvisited[graph.Vertices[i]] = true
	}
	return unvisited
}

func (graph *Graph) Length() int {
	return len(graph.Vertices)
}

func (graph *Graph) IncomingNodes() map[*Vertex]int {
	dict := make(map[*Vertex]int, 0)
	for _, v := range graph.Vertices {
		dict[v] = v.IncomingNodes
	}
	return dict
}

func (graph *Graph) createDistance() map[*Vertex]uint16 {
	dist := make(map[*Vertex]uint16, 0)
	for _, v := range graph.Vertices {
		dist[v] = math.MaxUint16
	}
	return dist
}
