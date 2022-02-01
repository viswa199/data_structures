package linkedlist

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type Single struct {
	Head   *Node
	length int
}

func (l *Single) Delete(data int){
	if l.Head==nil{
		fmt.Println("The list is empty")
	}else{
		if l.Head.Data==data{
			l.Head=l.Head.Next
		}else{
			temp:=l.Head
			for temp.Next!=nil{
				if temp.Next.Data==data{
					temp.Next=temp.Next.Next
					fmt.Println("Element Deleted")
					return
				}
				temp=temp.Next
			}
		}
	}
	fmt.Println("Sorry, Element not found")
}

func (l *Single) Insert(data int) {
	if l.Head == nil {
		l.Head = new(Node)
		l.Head.Data = data
		l.length++
	}else{
		temp:=l.Head
		l.Head=new(Node)
		l.Head.Data=data
		l.Head.Next=temp 
		l.length++
	}
}

func (l *Single) Search(a int) bool{
	for i:=l.Head; i!=nil ;{
		if i.Data==a{
			return true 
		}
		i=i.Next
	}
	return false
}

func (l *Single) Display() {
	for i := l.Head; i != nil; {
		fmt.Printf("%d ",i.Data)
		i=i.Next
	}
	fmt.Println()
}

func (l *Single) Len() int{
	return l.length
}
