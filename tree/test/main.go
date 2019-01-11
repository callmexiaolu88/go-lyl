package main

import (
	"fmt"
	"go-lyl/tree"
)

func main() {
	//arr := utils.RandFloat32Slice(1000, 20)
	arr := []float32{2, 10, 15, 3, 78, 52, 30, 7}
	root, err := tree.CreateTree(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Preorder")
	tree.PrintPreorder(root)
	fmt.Println()
	fmt.Println("Inorder")
	tree.PrintInorder(root)
	fmt.Println()
	fmt.Println("Postorder")
	tree.PrintPostorder(root)
	fmt.Println()
}
