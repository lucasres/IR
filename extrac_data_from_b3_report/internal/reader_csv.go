package internal

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/lucasres/IR/internal/dto"
)

type ReaderCsv struct{}

var reRealToFloar = regexp.MustCompile(`R\$\s*|\s*`)

func (r *ReaderCsv) GetData() ([]*dto.Entry, error) {
	f, err := os.Open("data.csv")
	if err != nil {
		return nil, fmt.Errorf("cant read data: %w", err)
	}

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cant read record from data: %w", err)
	}

	result := make([]*dto.Entry, 0)

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if row[2] != "Transferência - Liquidação" {
			continue
		}

		entry, err := r.parseRow(row)
		if err != nil {
			return nil, err
		}

		result = append(result, entry)
	}

	return result, nil
}

func (*ReaderCsv) parseRow(row []string) (*dto.Entry, error) {
	data, err := time.Parse("02/01/2006", row[1])
	if err != nil {
		return nil, fmt.Errorf("cant parse %s to time: %w", row[1], err)
	}

	movimentation := row[2]

	ticketSplited := strings.Split(row[3], "-")
	ticket := strings.Trim(ticketSplited[0], " ")

	amount, err := strconv.ParseInt(row[5], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("cant parse %s to int amount", row[5])
	}

	priceStr := strings.Trim(row[6], " ")
	priceStr = reRealToFloar.ReplaceAllString(row[6], "")
	priceStr = strings.ReplaceAll(priceStr, ",", ".")
	priceFl, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return nil, fmt.Errorf("cant parse %s to float price", row[6])
	}

	return &dto.Entry{
		Data:          data,
		Ticket:        ticket,
		Amount:        amount,
		Price:         int64(priceFl * 100),
		Movimentation: movimentation,
	}, nil
}
