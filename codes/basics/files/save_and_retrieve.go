package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func SaveAndRetrieve() {
	//data := "{\"employee\":{\"name\":\"Abhishek Ghosh\",\"salary\":10000,\"married\":false}}"
	//filename := "data.json"

	//saveToFile(filename, data)
	//fmt.Println(retrieveFromFile(filename))
	fmt.Println(appendToFile("data.txt", time.Now().String()))

}
func saveToFile(filename, data string) error {
	e := ioutil.WriteFile(filename, []byte(data), 0666)
	if e == nil {
		fmt.Println("data saved to file")
	} else {
		fmt.Println("data is not saved to file")
	}
	return e
}
func retrieveFromFile(filename string) string {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("unable to get the file", err)
		return ""
	}
	return string(bs)
}
func appendToFile(filename, newData string) string {

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(newData + "\n"); err != nil {
		panic(err)
	}

	data, err_ := ioutil.ReadFile(filename)
	if err_ != nil {
		fmt.Println(err)
	}
	return string(data)
}
