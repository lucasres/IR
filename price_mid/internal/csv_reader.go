package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type CsvReader struct{}

func (*CsvReader) ReadData() (map[Ticket][]Operation, error) {
	f, err := os.Open("data.csv")
	if err != nil {
		return nil, fmt.Errorf("cant read data: %w", err)
	}

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cant read record from data: %w", err)
	}

	result := make(map[Ticket][]Operation)
	for i, r := range records {
		if i == 0 {
			continue
		}

		t := Ticket(r[0])

		valueRaw := strings.Replace(r[2], ",", ".", -1)
		v, err := strconv.ParseFloat(valueRaw, 64)
		if err != nil {
			return nil, fmt.Errorf("cant parse %s to float: %w", valueRaw, err)
		}

		d, err := time.Parse("02/01/2006", r[4])
		if err != nil {
			return nil, fmt.Errorf("cant parse %s to time: %w", r[4], err)
		}

		a, err := strconv.Atoi(r[1])
		if err != nil {
			return nil, fmt.Errorf("cant parse %s to int: %w", r[1], err)
		}

		if _, has := result[t]; !has {
			result[t] = make([]Operation, 0)
		}

		result[t] = append(result[t], Operation{
			Ticket: t,
			Value:  int(v * 100),
			Data:   d,
			Amount: a,
		})
	}

	return result, nil
}
