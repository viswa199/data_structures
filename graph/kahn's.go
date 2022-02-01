//1.create a map of *Vertex and int to keep track of no of incoming edges to a node and intialize a map of *Vertex and bool to keep track of
//whether the node is visited or not.
//2.make of queue of the set of nodes of indegree zero and map the vertices present in the queue as true in visited.
//3.remove a vertex from the queue then:
//3a)==>reduce the indegree by 1 for all nodes adjacent to it.
//3b)==>if in-degree of an adjacent node is reduced to zero, then add it to the queue.
//4.repeat step 3 until the queue is empty.
//5.if the count of visited nodes is not equal to the nodes in the graph the the topological sort is not possible for the given graph.

package graph

import (
	"DS/queue"
)

func (graph *Graph) Kahn() (ordering []interface{}) {
	incomingnodes := graph.IncomingNodes()

	//a map of *Vertex and bool to track whether the node is visited.
	visited := graph.createVisited()

	//initializing an empty queue
	var q queue.Queue

	//adding elements to the queue if incomingnodes of a Vertex is zero.
	for vertice, inodes := range incomingnodes {
		if inodes == 0 {
			q.Enqueue(vertice)
			visited[vertice] = true
		}
	}

	for !q.Empty() {
		//remove an vertex from the queue.
		node := q.Dequeue().(*Vertex)
		ordering = append(ordering, node.Key)

		//reduce the indegree by 1 for all adjacent nodes to node vertex.
		for keys := range node.Adjacent {
			incomingnodes[keys] -= 1
		}

		//if indegree of an adjacent node is reduced to zero, then add it to the queue.
		for vertice, inodes := range incomingnodes {
			if inodes == 0 && !visited[vertice] {
				q.Enqueue(vertice)
				visited[vertice] = true
			}
		}
	}
	return
}
