package hashing

import "fmt"

type Node struct {
	N    int
	Next *Node
	Prev *Node
}

type Hash struct {
	Arr [10]*Node
}

func (h *Hash) Insert(a int) {
	hash_key := a % 10
	if h.Arr[hash_key] == nil {
		h.Arr[hash_key] = new(Node)
		h.Arr[hash_key].N = a
		h.Arr[hash_key].Next = nil
	} else {
		for h.Arr[hash_key].Next!= nil {
			h.Arr[hash_key] = h.Arr[hash_key].Next
		}
		h.Arr[hash_key].Next = new(Node)
		h.Arr[hash_key].Next.N = a
	}
}

func (h *Hash) Search(a int) bool {
	fmt.Println(h.Arr)
	hash_key := a % 10
	if h.Arr[hash_key] == nil {
		return false
	} else {
		temp := h.Arr[hash_key]
		for temp != nil {
			if temp.N == a {
				return true
			}
			temp = temp.Next
		}
	}
	return false
}

func (h *Hash) Delete(a int) (bool, int) {
	hash_key := a % 10
	if h.Arr[hash_key] == nil {
		return false, -1
	}
	temp := h.Arr[hash_key]
	fmt.Println(temp.Next.N)
	for temp != nil {
		if temp.N == a {
			temp =temp.Next
			h.Arr[hash_key]=temp
			return true, a
		}
		temp = temp.Next
	}
	return false, -1
}

func (h *Hash) Display() {
	for i := 0; i < 10; i++ {
		for h.Arr[i] != nil {
			fmt.Printf("%d ", h.Arr[i].N)
			h.Arr[i] = h.Arr[i].Next
		}
	}
}
