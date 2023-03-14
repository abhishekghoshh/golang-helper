package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

type city struct {
	name    string
	country string
}

func readCities(citiesChan chan []city) {
	var cities []city
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
	for _, line := range csvLines {
		cities = append(cities, city{
			name:    line[0],
			country: line[1],
		})
	}
	citiesChan <- cities
}

func createCity(record city, worker int) {
	time.Sleep(5 * time.Millisecond)
	f, err := os.OpenFile("cities-copy.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(record.name + "," + record.country + "," + strconv.Itoa(worker) + "\n"); err != nil {
		panic(err)
	}
}
func worker(cityChan chan city, worker int, jobStatus chan int) {
	for record := range cityChan {
		createCity(record, worker)
	}
	jobStatus <- worker
}
func main() {
	startTime := time.Now()

	citiesChan := make(chan []city)
	go readCities(citiesChan)

	const workers = 10
	jobs := make(chan city, 1600)
	jobStatus := make(chan int, workers)
	for w := 1; w <= workers; w++ {
		go worker(jobs, w, jobStatus)
	}
	fmt.Println("Trasferring data from one channel to different channel")
	counter := 0
	for _, cityChan := range <-citiesChan {
		counter++
		jobs <- cityChan
	}

	for cnt := 0; cnt <= 5; {
		select {
		case <-jobStatus:
			cnt++
		}
	}

	fmt.Println("\nrecord saved : ", counter)
	fmt.Println("total time elapsed ", time.Since(startTime))
}
