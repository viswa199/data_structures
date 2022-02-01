/*
1.Maintain a map of distance with *Vertex as key and int as value where the distance to every node is positive infinity. Mark the distance to the
start node 's' to be 0.
2.Maintain a PQ of key-value pairs of (*Vertex,distance) pairs which tell you which node to visit next based on sorted min value.
3.Insert (s,0) into the PQ and loop while PQ is not empty pulling out the next most promising (node,distance) pair.
4.Iterate over all edges outwards from the current node and relax each edge appending a new (node,distance) key-value pair to the PQ for every
relaxation.
*/

package graph

import (
	priorityqueue "DS/priority_queue"
	"log"
)

func (graph *Graph) Dijkstra(value interface{}) map[*Vertex]uint16 {

	visited := graph.createVisited()
	var pqueue priorityqueue.PQueue
	distance := graph.createDistance()

	node := graph.contains(value)
	if node == nil {
		log.Fatal("data doesn't exist.")
	}

	distance[node] = 0
	pqueue.Insert(&priorityqueue.Node{Value: node, Priority: distance[node]})
	visited[node] = true

	for !pqueue.Empty() {

		pNode := pqueue.Poll()
		node := pNode.Value.(*Vertex)

		for keys := range node.Adjacent {
			if pnode, ok := pqueue.Contains(keys); ok {
				distance[keys] = min(distance[keys], distance[node]+node.Adjacent[keys])
				pnode.Priority = min(pnode.Priority, distance[node]+node.Adjacent[keys])
			} else if !visited[keys] {
				distance[keys] = min(distance[keys], distance[node]+node.Adjacent[keys])
				pqueue.Insert(&priorityqueue.Node{Value: keys, Priority: distance[keys]})
			} else {
				continue
			}
		}

		pqueue.Delete()
		visited[node] = true
	}
	return distance
}

func (graph *Graph) ShortestPath(source, destination interface{}) {

}

func min(x, y uint16) uint16 {
	if x > y {
		return y
	}
	return x
}
