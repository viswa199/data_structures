package graph

import (
	"DS/stack"
	"fmt"
	"log"
)


func (graph *Graph) DepthFirstSearch(key interface{}) []interface{} {

	var s stack.Stack
	visited := graph.createVisited()

	var traversal []interface{}

	node := graph.contains(key)
	if node == nil {
		log.Fatal("key doesn't exist.")
	}

	s.Push(node)
	visited[node] = true
	traversal = append(traversal, node.Key)

	for !s.Empty() {
		node := s.Pop().(*Vertex)
		for keys := range node.Adjacent {
			if !s.Contains(keys) {
				if !visited[keys] {
					visited[keys] = true
					s.Push(keys)
					traversal = append(traversal, keys.Key)
					break
				}
			}
		}
	}
	return traversal
}

func (graph *Graph) RecursiveDFS(v interface{},visited map[*Vertex]bool){
	node:=graph.contains(v)
	if node==nil{
		log.Fatal("data doesn't exist.")
	}
	fmt.Printf("%v ",node.Key)
	visited[node]=true 
	for keys := range node.Adjacent {if _,ok:=visited[keys];!ok{
			graph.RecursiveDFS(keys.Key,visited)
		}
	}
}
