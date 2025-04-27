package internal

import (
	"fmt"
	"os"

	"github.com/lucasres/IR/internal/dto"
)

type WriterCSV struct{}

func (*WriterCSV) Save(entries []*dto.Entry) error {
	file, err := os.OpenFile("output.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	for _, entry := range entries {
		_, err := file.WriteString(entry.ToCsvRow() + "\n")
		if err != nil {
			return fmt.Errorf("error when write line: %s", entry.ToCsvRow())
		}
	}

	return nil
}
