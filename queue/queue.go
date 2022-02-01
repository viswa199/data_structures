package queue

import "fmt"

type Queue struct {
	t []interface{}
}

func (q *Queue) Enqueue(key interface{}) {
	q.t = append(q.t, key)
}

func (q *Queue) Dequeue() interface{} {
	deleting_element := q.t[0]
	q.t = q.t[1:]
	return deleting_element
}

func (q *Queue) Display() {
	for _, v := range q.t {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}

func (q *Queue) Length() int {
	return len(q.t)
}

func (q *Queue) Contains(key interface{}) bool {
	for _, v := range q.t {
		if v == key {
			return true
		}
	}
	return false
}

func (q *Queue) Empty() bool {
	if len(q.t) == 0 {
		return true
	}
	return false
}

func (q *Queue) Front() interface{} {
	return q.t[0]
}
