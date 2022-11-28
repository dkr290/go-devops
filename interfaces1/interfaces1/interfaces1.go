package main

import "fmt"

type Node interface {
	SetValue(i int)
	GetValue() int
}

type SLLNode struct {
	next  *SLLNode
	value int
}

func (s *SLLNode) SetValue(i int) {

	s.value = i
}

func (s *SLLNode) GetValue() int {

	return s.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

type PowerNode struct {
	next  *PowerNode
	value int
}

func (p *PowerNode) SetValue(i int) {

	p.value = i * 10
}

func (p *PowerNode) GetValue() int {

	return p.value
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

func intops() {

	var n Node
	n = NewSLLNode()
	n.SetValue(4)
	fmt.Println("Node is of value", n.GetValue())

	n = NewPowerNode()
	n.SetValue(2)
	fmt.Println("Node of value", n.GetValue())

	// if nn, ok := n.(*PowerNode); ok {
	// 	fmt.Println("This is the power node of", nn.value)
	// }

}
