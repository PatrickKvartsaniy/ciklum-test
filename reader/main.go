package main

import (
	"bufio"
	"ciklum/reader/models"
	"ciklum/reader/tools"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.csv")
	tools.CheckErr(err)
	defer file.Close()
	reader := csv.NewReader(bufio.NewReader(file))

	db := tools.CreateEngine()
	defer db.Close()
	tools.MakeMigrations()
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		customer := models.Customer{
			Name:  line[1],
			Email: line[2],
			Phone: line[3],
		}

		db.Create(&customer)
		fmt.Printf("Customer saved\n")
	}
}
