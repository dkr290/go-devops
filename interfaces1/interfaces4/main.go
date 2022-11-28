package main

import (
	"fmt"

	"github.com/dkr290/go-devops/interfaces1/interfaces4/powernode"
	"github.com/dkr290/go-devops/interfaces1/interfaces4/sllnode"
)

type Node interface {
	SetValue(i int) error
	GetValue() int
}

func main() {

	n := createNode(5)

	switch value := n.(type) {
	case *sllnode.SLLNode:
		fmt.Println("Type is SLLNode, message:", value.SnodeMessage, value.GetValue())

	case *powernode.PowerNode:
		fmt.Println("Type is PowerNode, message:", value.PnodeMessage, value.GetValue())

	}
}

func createNode(v int) Node {
	sn := sllnode.NewSLLNode()
	sn.SetValue(v)
	return sn
}
