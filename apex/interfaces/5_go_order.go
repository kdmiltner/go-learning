package main

import (
	"errors"
	"fmt"
	"log"
)

type Processor interface {
	Process() (string, error)
}

type OrderRequestDNA struct {
	ID                string
	FulfillmentMethod string
	TotalItems        uint32
	TotalPaid         float32
	Partner           string
}

type OrderRequestDXE struct {
	DXEID       string
	CFAID       string
	Fulfillment string
	TenderAmt   float32
	Partner     string
}

func (o *OrderRequestDNA) Process() (string, error) {
	empty := OrderRequestDNA{}
	if *o == empty {
		return "", errors.New("empty OrderRequestDNA")
	}
	return "logic to process Order DNA flow", nil
}

func (o *OrderRequestDXE) Process() (string, error) {
	empty := OrderRequestDXE{}
	if *o == empty {
		return "", errors.New("empty OrderRequestDXE")
	}
	return "logic to process DXE Order flow", nil
}

func CreateOrder(p Processor) (string, error) {
	if p == nil {
		return "", errors.New("nil Processor")
	}
	return p.Process()
}

func main() {
	featureFlag := "OrderDNA"
	var req Processor

	switch featureFlag {
	case "OrderDNA":
		req = &OrderRequestDNA{
			ID:                "123",
			FulfillmentMethod: "fulfilled by UberEats",
			TotalItems:        4,
			TotalPaid:         16.50,
			Partner:           "UberEats",
		}
	case "OrderDXE":
		req = &OrderRequestDXE{
			DXEID:       "123",
			CFAID:       "CFAID-1234",
			Fulfillment: "fulfilled by UberEats",
			TenderAmt:   16.50,
			Partner:     "UberEats",
		}
	}

	o, err := CreateOrder(req) // One way to do this
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(o)

	//o, err := req.Process() // Another way to do this
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(o)
}
