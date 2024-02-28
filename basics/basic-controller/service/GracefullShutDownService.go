package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
)

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		// write your own cleanup function
		saveToJson()
		os.Exit(0)
	}()
}

func saveToJson() {
	data, err := json.Marshal(GetPersons())
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("persons.json", []byte(data), 0666)
}
