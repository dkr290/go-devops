package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address1 struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zipcode int    `json:"zipcode"`
}

type Item struct {
	Name        string `json:"itemname"`
	Description string `json:"desc,omitempty"`
	Quantity    int    `json:"qty"`
	Price       int    `json:"price"`
}

type Order struct {
	TotalPrice  int    `json:"totalprice,omitempty"`
	IsPayed     bool   `json:"payed"`
	Fragile     bool   `json:"fragile,omitempty"`
	OrderDetail []Item `json:"orderdetail"`
}

type Customer struct {
	UserName      string   `json:"username"`
	Password      string   `json:"-"`
	Token         string   `json:"-"`
	ShipTo        Address1 `json:"shipto"`
	PurchaseOrder Order
}

func jsonDataDecode(b []byte) *Customer {

	if !json.Valid(b) {
		fmt.Println("The json data is not valid", string(b))
		os.Exit(1)
	}

	var c Customer

	err := json.Unmarshal(b, &c)
	if err != nil {
		log.Fatalln("Cannot unmarshal the data", err)
	}

	return &c

}

func (c *Customer) Total() {

	price := 0
	for _, item := range c.PurchaseOrder.OrderDetail {
		price += item.Quantity * item.Price
	}
	c.PurchaseOrder.TotalPrice = price

}
