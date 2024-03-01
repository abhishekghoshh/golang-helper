package concurrenypattern

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

type City struct {
	name    string
	country string
}

const FILE_SRC = "concurrenypattern/resources/all-city.csv"
const FILE_DEST = "concurrenypattern/resources/all-city-copy.csv"
const FILE_DEST_FOR_WORKER_POOL = "concurrenypattern/resources/all-city-copy-worker-pool.csv"

func watchFunction(identifier string, runnable func()) {
	startTime := time.Now()
	runnable()
	fmt.Println("total time elapsed for", identifier, time.Since(startTime))
}

func CusomWorkerPoolExample1() {
	watchFunction("without-wokrer-pool", withoutWorkerPool)
	watchFunction("with-wokrer-pool", withWorkerPool)
}

func withoutWorkerPool() {
	cityCsv, err := os.Open(FILE_SRC)
	if err != nil {
		panic(err)
	}
	defer cityCsv.Close()
	fmt.Println("Successfully opened", FILE_SRC)

	csvLines, err := csv.NewReader(cityCsv).ReadAll()
	if err != nil {
		panic(err)
	}

	counter := 0
	for _, line := range csvLines {
		counter++
		createCity(
			City{
				name:    line[0],
				country: line[1],
			})
	}
	fmt.Println("record saved : ", counter)
}
func createCity(record City) {
	f, err := os.OpenFile(FILE_DEST, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(record.name + "," + record.country + "\n"); err != nil {
		panic(err)
	}
}

// TODO unable to read from the file properly beacuse
// a single text file can not be read in chunks as the we can not directly jump to the nth line
// we have to traverse the entire other
// the concurrency pattern is perfect
func withWorkerPool() {
	const workers = 2000
	const lineLength = 44691
	const chunkSize = lineLength / workers
	cityChan := make(chan City, lineLength)
	workerChan := make(chan bool, workers)
	for w := 0; w < workers; w++ {
		file, err := os.Open(FILE_SRC)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		start := chunkSize * w
		end := chunkSize * (w + 1)
		if end >= lineLength {
			end = lineLength - 1
		}
		go func(s, e int64) {
			reader := csv.NewReader(file)
			file.Seek(s, 0)
			for i := s; i < e; i++ {
				record, err := reader.Read()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(i, " record is", record)
				row := City{
					name:    record[0],
					country: record[1],
				}
				cityChan <- row
			}
			workerChan <- true
		}(int64(start), int64(end))
	}
	go func() {
		counter := 0
		for range workerChan {
			counter++
			if counter == workers {
				close(cityChan)
				close(workerChan)
				return
			}
		}
	}()
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			f, err := os.OpenFile(FILE_DEST_FOR_WORKER_POOL, os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			for record := range cityChan {
				if _, err = f.WriteString(record.name + "," + record.country + "," + strconv.Itoa(w) + "\n"); err != nil {
					panic(err)
				}
			}
		}()
	}
	wg.Wait()
}
