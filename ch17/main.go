package main

import (
	"fmt"
	"time"
)

type Courier struct {
	Name string
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	p := &Parcel{}
	p.Pdt = pdt
	p.ShippedTime = time.Now()
	return p
}

type Product struct {
	Name  string
	Price int
	ID    int
}

type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeliveredTime time.Time
}

func (p *Parcel) Delivered() *Product {
	p.DeliveredTime = time.Now()
	return p.Pdt
}

func main() {
	c := Courier{"택배회사"}
	pdt := Product{"물품", 10000, 1}
	p := c.SendProduct(&pdt)
	fmt.Println("p=", p)

	p.Delivered()
	fmt.Println("p=", p)
}
