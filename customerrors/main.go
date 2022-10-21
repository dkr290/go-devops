package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() {

	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidRoutingNumber {
				fmt.Printf("invalid routing number %d and is < 100\n", d.routingNumber)
			}
		}
	}()

	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}

}

func (d *directDeposit) validateLastName() error {

	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	fmt.Println("Last Name:", d.lastName)
	fmt.Println("First Name:", d.firstName)
	fmt.Println("Bank Name:", d.bankName)
	fmt.Println("Routing Number:", d.routingNumber)
	fmt.Println("Account Number:", d.accountNumber)
}

func main() {

	d1 := directDeposit{
		firstName:     "James",
		bankName:      "XYZ inc",
		routingNumber: 17,
		accountNumber: 1809,
	}

	d1.validateRoutingNumber()

	err := d1.validateLastName()
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("===========================================================================")
	d1.report()
}
