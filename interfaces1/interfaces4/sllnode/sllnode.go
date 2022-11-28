package sllnode

import "github.com/dkr290/go-devops/interfaces1/interfaces4/customerrors"

type SLLNode struct {
	next         *SLLNode
	value        int
	SnodeMessage string
}

func (s *SLLNode) SetValue(i int) error {

	if s == nil {
		return customerrors.ErrInvalidNode
	}
	s.value = i
	return nil
}

func (s *SLLNode) GetValue() int {

	return s.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{SnodeMessage: "This is message from the normal node"}
}
