package tree

import (
	"errors"
	"fmt"
)

type CompareResult int

const (
	AGB CompareResult = iota
	AGEB
	ALB
	ALEB
	AEB
)

type Node struct {
	Data  float32
	Left  *Node
	Right *Node
}

var DataNotFindException = fmt.Errorf("Data not find")
var DataAlreadyExistException = fmt.Errorf("Data already exists")
var TreeNilException = fmt.Errorf("tree is nil")

func Insert(tree *Node, key float32) (bool, error) {
	if tree != nil {
		if tree.Data > key {
			if tree.Left != nil {
				return Insert(tree.Left, key)
			}
			tree.Left = &Node{Data: key}
			return true, nil
		} else if tree.Data < key {
			if tree.Right != nil {
				return Insert(tree.Right, key)
			}
			tree.Right = &Node{Data: key}
			return true, nil
		}
		return false, DataAlreadyExistException
	}
	return false, TreeNilException
}

func Delete(tree *Node, key float32) (bool, error) {
	if tree != nil {
		if tree.Data > key {
			if tree.Left != nil {
				return Delete(tree.Left, key)
			}
			return false, DataNotFindException
		} else if tree.Data < key {
			if tree.Right != nil {
				return Insert(tree.Right, key)
			}
			return false, DataNotFindException
		}
		//左节点是nil
		if tree.Left == nil {
			*tree = *tree.Right
		} else if tree.Right == nil {
			*tree = *tree.Left
		} else {
			t := tree.Right
			for t.Left != nil {
				t = t.Left
			}
			tree.Data = t.Data
			*t = *t.Right
		}
		return false, DataAlreadyExistException
	}
	return false, TreeNilException
}

func Search(tree *Node, key float32) (*Node, error) {
	if tree != nil {
		if tree.Data > key {
			if tree.Left != nil {
				return Search(tree.Left, key)
			}
			return nil, DataNotFindException
		} else if tree.Data < key {
			if tree.Right != nil {
				return Search(tree.Right, key)
			}
			return nil, DataNotFindException
		}
		return tree, nil
	}
	return nil, TreeNilException
}

func CreateTree(arr []float32) (*Node, error) {
	if len(arr) > 0 {
		tree := &Node{
			Data: arr[0],
		}
		for _, v := range arr[1:] {
			_, err := Insert(tree, v)
			if err != nil {
				return nil, err
			}
		}
		return tree, nil
	}
	return nil, errors.New("array is empty")
}

func PrintInorder(tree *Node) {
	if tree != nil {
		PrintInorder(tree.Left)
		printData(tree)
		PrintInorder(tree.Right)
	}
}

func PrintPreorder(tree *Node) {
	if tree != nil {
		printData(tree)
		PrintPreorder(tree.Left)
		PrintPreorder(tree.Right)
	}
}

func PrintPostorder(tree *Node) {
	if tree != nil {
		PrintPostorder(tree.Left)
		PrintPostorder(tree.Right)
		printData(tree)
	}
}

func printData(tree *Node) {
	if tree != nil {
		fmt.Print(tree.Data, " ")
	}
}
