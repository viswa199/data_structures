package queue

import "fmt"

type Queue struct {
	t []string
}

func (q *Queue) Enqueue(key string) {
	q.t = append(q.t, key)
}

func (q *Queue) Dequeue() string {
	deleting_element := q.t[len(q.t)-1]
	q.t = q.t[1:]
	return deleting_element
}

func (q *Queue) Display() {
	for _, v := range q.t {
		fmt.Printf("%s ",v)
	}
	fmt.Println()
}

func (q *Queue) Length() int{
	return len(q.t)
}

func (q *Queue) Contains(key string) bool{
	for _,v:=range q.t{
		if v==key{
			return true
		}
	}
	return false
}

func (q *Queue) Empty() bool{
	if len(q.t)==0{
		return true
	}
	return false
}

func (q *Queue) Front() string{
	return q.t[0]
}