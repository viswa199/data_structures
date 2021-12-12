package trees

import "fmt"

type BST struct {
	Root *Node
	Len  int
}

func (t *BST) Add(data int) {
	t.Insert(data, t.Root)
	t.Len++
}

func (t *BST) Insert(data int, root *Node) *Node {

	if t.Root == nil {
		t.Root = &Node{Data: data}
		return t.Root
	}

	if root == nil {
		return &Node{Data: data}
	} else {
		if root.Data < data {
			root.Right = t.Insert(data, root.Right)
		} else {
			root.Left = t.Insert(data, root.Left)
		}
	}
	return root
}

func (t *BST) Display(root *Node) {
	if root == nil {
		return
	}
	t.Display(root.Left)
	fmt.Println(root.Data)
	t.Display(root.Right)
}

func (t *BST) Search(root *Node, data int) (*Node, bool) {
	if root == nil {
		return nil, false
	} else {
		if root.Data == data {
			return root, true
		} else if root.Data > data {
			return t.Search(root.Left, data)
		} else {
			return t.Search(root.Right, data)
		}
	}
}

func (t *BST) Remove(root *Node, data int) *Node {
	return nil
}

func (t *BST) Length() int {
	return t.Len
}

func (t *BST) InorderPredecessor(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	if root.Right == nil {
		return root.Left
	}
	return t.InorderPredecessor(root.Left)
}

func (t *BST) InorderSuccessor(root *Node) *Node {
	if root == nil {
		return nil
	}
	if (root.Left == nil && root.Right == nil) || root.Left == nil {
		return root
	}
	return t.InorderSuccessor(root.Right)
}

func (t *BST) Delete(root *Node, val int) *Node {
	if nil == root {
		return root
	}
	if root.Data > val {
		root.Left = t.Delete(root.Left, val)
	}
	if root.Data < val {
		root.Right = t.Delete(root.Right, val)
	}
	if root.Data == val {
		if root.Left == nil && root.Right == nil {
			root = nil
			return root
		}
		if root.Left == nil && root.Right != nil {
			temp := root.Right
			root = nil
			root = temp
			return root
		}
		if root.Left != nil && root.Right == nil {
			temp := root.Left
			root = nil
			root = temp
			return root
		}

		tempNode := minValued(root.Right)
		root.Data = tempNode.Data
		root.Right = t.Delete(root.Right, tempNode.Data)
	}
	return root
}

func minValued(root *Node) *Node {
	temp := root
	for nil != temp && temp.Left != nil {
		temp = temp.Left
	}
	return temp
}

func (t *BST) IsBST(root *Node) bool{
	if root==nil{
		return true 
	}
	if root.Left!=nil{
		if root.Data<=root.Left.Data{
			return false 
		}
	}
	if root.Right!=nil{
		if root.Data>=root.Right.Data{
			return false 
		}
	}
	if t.IsBST(root.Left) && t.IsBST(root.Right) {
		return true 
	}
	return false
}
