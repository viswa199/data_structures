package priorityqueue

import (
	"fmt"
	"log"
)

type PQueue struct {
	Arr []*Node
}

type Node struct {
	Value    interface{}
	Priority uint16
}

func (p *PQueue) Insert(node *Node) {
	p.Arr = append(p.Arr, node)
}

func (p *PQueue) maxheap() {
	l := len(p.Arr)
	i := l / 2
	for j := i - 1; j >= 0; j-- {
		if 2*j+1 < l {
			if p.Arr[2*j+1].Priority > p.Arr[j].Priority {
				p.Arr[2*j+1], p.Arr[j] = p.Arr[j], p.Arr[2*j+1]
			}
		}
		if 2*j+2 < l {
			if p.Arr[2*j+2].Priority > p.Arr[j].Priority {
				p.Arr[2*j+2], p.Arr[j] = p.Arr[j], p.Arr[2*j+2]
			}
		}
	}
}

func (p *PQueue) Delete() interface{} {
	l := len(p.Arr)
	if l == 0 {
		log.Fatal("Trying to delete an element from an empty slice.")
	}
	item := p.Arr[l-1]
	p.Arr = p.Arr[1:]
	return item.Value
}

func (p PQueue) Print() {
	for _, v := range p.Arr {
		fmt.Printf("%v ", v.Value)
	}
	fmt.Println()
}

func (p *PQueue) minHeap() {
	l := len(p.Arr)
	no_of_roots := l/2 - 1
	for i := no_of_roots; i >= 0; i-- {
		if 2*i+1 < l {
			if p.Arr[2*i+1].Priority < p.Arr[i].Priority {
				p.Arr[2*i+1], p.Arr[i] = p.Arr[i], p.Arr[2*i+1]
			}
		}
		if 2*i+2 < l {
			if p.Arr[2*i+2].Priority < p.Arr[i].Priority {
				p.Arr[2*i+2], p.Arr[i] = p.Arr[i], p.Arr[2*i+2]
			}
		}
	}
}

func (p *PQueue) Remove(value interface{}) {
	for i, v := range p.Arr {
		if v.Value == value {
			p.Arr = append(p.Arr[:i], p.Arr[i+1:]...)
		}
	}
}

func (p *PQueue) Poll() *Node{
	p.minHeap()
	return p.Arr[0]
}

func (p *PQueue) Contains(value interface{}) (*Node, bool) {
	for _, v := range p.Arr {
		if v.Value == value {
			return v,true 
		}
	}
	return nil,false 
}

func (p *PQueue) Empty() bool {
	if len(p.Arr) == 0 {
		return true
	}
	return false
}
