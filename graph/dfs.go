package graph

import "DS/queue"

func Depth_First_Traversal(g Graph) []string{
	var visited []string
	q := queue.Queue{}
	q.Enqueue("A")
	for !q.Empty() || len(visited) == len(g.Vertices) {
		visited = append(visited, q.Front())
		for _, v := range g.getVertex(q.Front()).Adjacent {
			if !q.Contains(v.Key) {
				q.Enqueue(v.Key)
			}
		}
		q.Dequeue()
	}
	return visited
}
