package main

import (
	"fmt"

	"github.com/lucasres/calc_ir/internal"
)

func main() {
	reader := &internal.CsvReader{}

	data, err := readData(reader)
	if err != nil {
		panic(err)
	}

	for ticket, operations := range data {
		totalAmount := 0
		totalValue := 0

		for _, o := range operations {
			totalAmount += o.Amount
			totalValue += o.Value * o.Amount
		}

		midPrice := (totalValue / totalAmount)

		fmt.Printf("Ticket: %s -> %d \n", ticket, midPrice)
	}

}

func readData(r internal.Reader) (map[internal.Ticket][]internal.Operation, error) {
	return r.ReadData()
}
