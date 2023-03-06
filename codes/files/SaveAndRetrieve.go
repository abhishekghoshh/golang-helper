package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//data := "{\"employee\":{\"name\":\"Abhishek Ghosh\",\"salary\":10000,\"married\":false}}"
	filename := "data.json"

	//saveToFile(filename, data)
	fmt.Println(retrieveFromFile(filename))

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
