package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"
)

// Processor is implemented by Process and can be used to process various types of data.
type Processor interface {
	Process(bool) error
}

// Order represents an order.
type Order struct {
	ID             string `json:"id"`
	Partner        string `json:"partner"`
	LocationNumber string `json:"locationNumber"`
}

// Response represents the response received when creating an order.
type Response struct {
	ID        string
	CreatedAt time.Time
}

// GetOrder takes the []byte provided and unmarshals it to an Order struct.
// Returns a Processor.
func GetOrder(b []byte) (Processor, error) {
	order := &Order{}
	empty := Order{}

	err := json.Unmarshal(b, order)
	if *order == empty {
		return nil, errors.Join(errors.New("empty order"), err)
	}

	return order, err
}

// CreateOrder creates an order using the id provided. If fail is set to
// true it will always respond with an empty Response and an error.
// If fail is set to false it will always respond with a Response and nil.
func CreateOrder(id string, fail bool) (Response, error) {
	fmt.Println("calling CreateOrder")
	if fail {
		return Response{}, errors.New("CreateOrder failed")
	}

	return Response{ID: "1234", CreatedAt: time.Now()}, nil
}

// Process will call CreateOrder, save to the database, and calls back to UberEats
// to accept the order.
func (o *Order) Process(fail bool) error {
	resp, err := CreateOrder(o.ID, fail)
	if err != nil {
		return errors.Join(errors.New("process order failed"), err)
	}

	fmt.Printf("save resp to database - resp: %#v\n", resp)
	fmt.Println("call UE accept order endpoint")

	return nil
}

func main() {
	// Mock receiving a JSON object.
	b := []byte("{\"id\":\"1234\",\"partner\":\"UberEats\",\"locationNumber\":\"00440\"}")

	// Call GetOrder, which will unmarshal the JSON object into the Order struct.
	o, err := GetOrder(b)
	// Handle error.
	if err != nil {
		panic(err) // panic because if we don't have an order we can't do anything else.
	}

	// Call o.Process(true) will return an error.
	// Call o.Process(false) will return success.
	if err := o.Process(false); err != nil {
		log.Fatal(err)
	}
}
