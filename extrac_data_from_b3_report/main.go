package main

import (
	"log"

	"github.com/lucasres/IR/internal"
)

func main() {
	reader := &internal.ReaderCsv{}
	writer := &internal.WriterCSV{}

	entries, err := reader.GetData()
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Save(entries)
	if err != nil {
		log.Fatal(err)
	}
}
