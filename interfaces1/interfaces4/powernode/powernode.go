package powernode

import "github.com/dkr290/go-devops/interfaces1/interfaces4/customerrors"

type PowerNode struct {
	next         *PowerNode
	value        int
	PnodeMessage string
}

func (p *PowerNode) SetValue(v int) error {
	if p == nil {
		return customerrors.ErrInvalidNode
	}

	p.value = v * 10
	return nil
}

func (p *PowerNode) GetValue() int {

	return p.value

}

func NewPowerNode() *PowerNode {
	return &PowerNode{PnodeMessage: "This is the message from the powernode"}
}
