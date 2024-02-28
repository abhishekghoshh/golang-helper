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

func readCities(cityListChan chan<- []city) {
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
	cityListChan <- cities
}

func createCity(record city, worker int) {
	time.Sleep(2 * time.Millisecond)
	f, err := os.OpenFile("cities-copy.csv", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(record.name + "," + record.country + "," + strconv.Itoa(worker) + "\n"); err != nil {
		panic(err)
	}
}
func worker(cityChan <-chan city, worker int, jobChan chan bool) {
	for record := range cityChan {
		createCity(record, worker)
	}
	jobChan <- true
}
func main() {
	startTime := time.Now()

	cityListChan := make(chan []city)
	go readCities(cityListChan)

	const workers = 20
	cityChan := make(chan city, 1600)
	jobChan := make(chan bool, workers)

	for w := 1; w <= workers; w++ {
		go worker(cityChan, w, jobChan)
	}
	fmt.Println("Trasferring data from one channel to different channel")
	counter := 0
	for _, oneCity := range <-cityListChan {
		counter++
		cityChan <- oneCity
	}

	// blocking call to check if the previous go routines are done or not
	isClosed := false
	for true {
		if len(cityChan) == 0 && !isClosed {
			close(cityChan)
			isClosed = true
		}
		if len(jobChan) == workers {
			close(jobChan)
			break
		}
	}

	fmt.Println("record saved : ", counter)
	fmt.Println("total time elapsed ", time.Since(startTime))
}
