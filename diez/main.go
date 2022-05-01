package main

import (
	"fmt"
	"vasquezruiz.com/datastructs"
)

func main() {
	fmt.Println("*** Data Structs ***")
	var intTree datastructs.BinaryTree[int]

	fmt.Printf("Int Tree %v\n", intTree)

	initTree := datastructs.NewBinaryTree(10)
	datastructs.Insert(initTree, 5)
	datastructs.Insert(initTree, 20)
	datastructs.Insert(initTree, 2)
	fmt.Println("Numbers tree")
	datastructs.Inorder(initTree)

	strTree := datastructs.NewBinaryTree("hola")
	datastructs.Insert(strTree, "borola")
	datastructs.Insert(strTree, "nariz")
	datastructs.Insert(strTree, "de")
	datastructs.Insert(strTree, "bola")
	fmt.Println("String tree")
	datastructs.Inorder(strTree)
}