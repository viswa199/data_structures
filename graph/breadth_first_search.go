package graph

import (
	"DS/queue"
	"fmt"
	"log"
)

func (graph *Graph) BreadthFirstSearch(Key interface{}) []interface{} {
	var q queue.Queue
	visited:=graph.createVisited()
	var traversal []interface{}

	node:=graph.contains(Key)
	if node==nil{
		log.Fatal("data doesn't exist")
	}
	q.Enqueue(node)
	visited[node]=true 
	traversal = append(traversal, node.Key)

	for !q.Empty(){
		node:=q.Dequeue().(*Vertex)
		fmt.Println(node.Key)
		for keys := range node.Adjacent {	if !q.Contains(keys){
				if !visited[keys]{
					q.Enqueue(keys)
					visited[keys]=true 
					traversal = append(traversal, keys.Key)
				}
			}
		}
	}
	return traversal
}