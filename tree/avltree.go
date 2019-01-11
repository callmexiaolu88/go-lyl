package tree

type AVLNode struct {
	BFSeek int
	Data   float32
	Left   *AVLNode
	Right  *AVLNode
}

func AVLInsert(tree *AVLNode, key float32) (bool, error) {
	if tree != nil {
		if tree.Data > key {
			if tree.Left != nil {
				return AVLInsert(tree.Left, key)
			}
			tree.Left = &AVLNode{Data: key}

			return true, nil
		} else if tree.Data < key {
			if tree.Right != nil {
				return AVLInsert(tree.Right, key)
			}
			tree.Right = &AVLNode{Data: key}
			return true, nil
		}
		return false, DataAlreadyExistException
	}
	return false, TreeNilException
}
