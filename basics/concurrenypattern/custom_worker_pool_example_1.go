package concurrenypattern

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const FILE_SRC = "concurrenypattern/resources/customers-500000.csv"
const FILE_DEST = "concurrenypattern/resources/copy.csv"
const FILE_DEST_FOR_WORKER_POOL = "concurrenypattern/resources/copy-with-workerpool.csv"

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
	fileSrc, err := os.Open(FILE_SRC)
	if err != nil {
		panic(err)
	}
	defer fileSrc.Close()

	fileDest, err := os.OpenFile(FILE_DEST, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer fileDest.Close()

	scanner := bufio.NewScanner(fileSrc)
	for scanner.Scan() {
		data := scanner.Text()
		if _, err = fileDest.WriteString(data + "\n"); err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// TODO unable to read from the file properly beacuse
// a single text file can not be read in chunks as the we can not directly jump to the nth line
// we have to traverse the entire other
// the concurrency pattern is perfect
func withWorkerPool() {
	capacity := 100000
	worker := 1000

	ch := make(chan string, capacity)
	marker := make(chan bool)

	go func() {
		for {
			select {
			case <-marker:
				close(ch)
			default:
				continue
			}
		}
	}()

	file, err := os.Open(FILE_SRC)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < worker; i++ {
		go saveToFile(ch)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	marker <- true

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func saveToFile(ch chan string) {
	f, err := os.OpenFile(FILE_DEST_FOR_WORKER_POOL, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for data := range ch {
		if _, err = f.WriteString(data + "\n"); err != nil {
			panic(err)
		}
	}
}
