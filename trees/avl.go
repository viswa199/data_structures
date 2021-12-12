package trees

import "fmt"

type AVLTree struct {
	root *AVLNode
}

func (t *AVLTree) Add(value int) {
	t.root = t.root.add(value)
}

func (t *AVLTree) Remove(value int) {
	t.root = t.root.remove(value)
}

func (t *AVLTree) Update(olderValue, newValue int) {
	t.root = t.root.remove(olderValue)
	t.root = t.root.add(newValue)
}

func (t *AVLTree) Search(value int) (node *AVLNode) {
	return t.root.search(value)
}

func (t *AVLTree) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

func (t *AVLTree) DisplayPreOrder(){
	t.root.displayPreorder()
}

// AVLNode structure
type AVLNode struct {
	value int

	// height counts nodes (not edges)
	height int
	left   *AVLNode
	right  *AVLNode
}

// Adds a new node
func (n *AVLNode) add(value int) *AVLNode {
	if n == nil {
		return &AVLNode{value, 1, nil, nil}
	}

	if value < n.value {
		n.left = n.left.add(value)
	} else if value > n.value {
		n.right = n.right.add(value)
	} else {
		// if same key exists update value
		n.value = value
	}
	return n.rebalanceTree()
}

// Removes a node
func (n *AVLNode) remove(key int) *AVLNode {
	if n == nil {
		return nil
	}
	if key < n.value {
		n.left = n.left.remove(key)
	} else if key > n.value {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := n.right.findSmallest()
			n.value = rightMinNode.value
			n.value = rightMinNode.value
			// delete smallest node that we replaced
			n.right = n.right.remove(rightMinNode.value)
		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

// Searches for a node
func (n *AVLNode) search(value int) *AVLNode {
	if n == nil {
		return nil
	}
	if value < n.value {
		return n.left.search(value)
	} else if value > n.value {
		return n.right.search(value)
	} else {
		return n
	}
}

// Displays nodes left-depth first (used for debugging)
func (n *AVLNode) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.value, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

func (n *AVLNode) displayPreorder(){
	if n==nil{
		return
	}
	fmt.Printf("%d ",n.value)
	n.left.displayPreorder()
	n.right.displayPreorder()
}

// Checks if node is balanced and rebalance
func (n *AVLNode) rebalanceTree() *AVLNode {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *AVLNode) rotateLeft() *AVLNode {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *AVLNode) rotateRight() *AVLNode {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Finds the smallest child (based on the key) for the current node
func (n *AVLNode) findSmallest() *AVLNode {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

// Returns max number - TODO: std lib seemed to only have a method for floats!
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
