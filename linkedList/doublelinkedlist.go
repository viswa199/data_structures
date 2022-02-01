package linkedlist

import "fmt"

type Node1 struct {
	Data int
	Next *Node1
	Prev *Node1
}

type Double struct {
	Head   *Node1
	Tail   *Node1
	Length int
}

func (l *Double) Insert(data int) {
	if l.Head == nil {
		l.Head = new(Node1)
		l.Tail = l.Head
		l.Head.Data = data
		l.Length++
	} else {
		temp := l.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = new(Node1)
		temp.Next.Data = data
		temp.Next.Prev = temp
		l.Tail = temp.Next
		l.Length++
	}
}

func (l *Double) ReverseDisplay() {
	for i := l.Tail; i != nil; {
		fmt.Printf("%d ",i.Data)
		i=i.Prev
	}
	fmt.Println()
}

func (l *Double) Display(){
	for i:=l.Head; i!=nil;{
		fmt.Printf("%d ",i.Data)
		i=i.Next
	}
	fmt.Println()
}

func (l *Double) Len() int{
	return l.Length
}

func (l *Double) Reverse(){
	i:=l.Head
	j:=l.Tail
	for k:=0;k<(l.Length)/2;k++{
		if i!=nil && j!=nil{
			i.Data,j.Data=j.Data,i.Data
			i=i.Next
			j=j.Prev
		}
	}
}

func (l *Double) Delete(x int) (bool,error){
	if l.Head==nil{
		return false,fmt.Errorf("List is empty")
	}else{
		if l.Head.Data==x{
			l.Head=l.Head.Next
			return true,nil 
		}else{
			temp:=l.Head
			for temp.Next!=nil{
				if temp.Next.Data==x{
					temp.Next=temp.Next.Next
					return true,nil 
				}
				temp=temp.Next
			}
		}
	}
	return false,fmt.Errorf("The element %d is not present in list",x)
}