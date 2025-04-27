package dto

import (
	"fmt"
	"time"
)

type Entry struct {
	Ticket        string
	Data          time.Time
	Amount        int64
	Price         int64
	Movimentation string
}

func (e *Entry) ToCsvRow() string {
	return fmt.Sprintf(
		"%s,%d,%d,%s,%s,%f",
		e.Ticket,
		e.Amount,
		e.Price,
		"Compra",
		e.Data.Format("02/01/2006"),
		float64((e.Amount*e.Price))/100,
	)
}
