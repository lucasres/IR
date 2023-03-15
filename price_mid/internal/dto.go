package internal

import "time"

type Ticket string

type Operation struct {
	Ticket Ticket
	// int onde os 2 ultimos são os centavos
	Value  int
	Amount int
	Data   time.Time
}

type Reader interface {
	ReadData() (map[Ticket][]Operation, error)
}
