package trees

import "fmt"

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
	Len  int
}

func (t *Tree) Insert(Data int) *Node {
	NewNode := new(Node)
	if t.Root == nil {
		t.Root = NewNode
		t.Root.Data = Data
		t.Root.Left = nil
		t.Root.Right = nil
	} else {
		NewNode.Data = Data
		NewNode.Left = nil
		NewNode.Right = nil
	}
	t.Len++
	return NewNode
}

func (t *Tree) Length() int{
	return t.Len
}

func (t *Tree) Preorder(root *Node) {
	if root==nil{
		return
	}
	fmt.Printf("%d ",root.Data)
	t.Preorder(root.Left)
	t.Preorder(root.Right)
}

func (t *Tree) Postorder(root *Node){
	if root==nil{
		return
	}
	t.Postorder(root.Left)
	t.Postorder(root.Right)
	fmt.Printf("%d ",root.Data)
}

func (t *Tree) Inorder(root *Node){
	if root==nil{
		return
	}
	t.Inorder(root.Left)
	fmt.Printf("%d ",root.Data)
	t.Inorder(root.Right)
}
