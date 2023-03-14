package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type city struct {
	name    string
	country string
}

func createCity(record city) {
	time.Sleep(5 * time.Millisecond)
	f, err := os.OpenFile("cities-copy.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(record.name + "," + record.country + "\n"); err != nil {
		panic(err)
	}
}

func main() {
	startTime := time.Now()
	cityCsv, err := os.Open("cities.csv")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Successfully opened cities.csv")
	defer cityCsv.Close()

	csvLines, err := csv.NewReader(cityCsv).ReadAll()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	counter := 0
	for _, line := range csvLines {
		counter++
		createCity(city{
			name:    line[0],
			country: line[1],
		})
	}

	fmt.Println("record saved : ", counter)
	fmt.Println("total time elapsed ", time.Since(startTime))
}
