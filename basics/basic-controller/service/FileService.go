package service

import (
	"fmt"
	"io/ioutil"
)

func SaveToFile(filename, data string) error {
	e := ioutil.WriteFile(filename, []byte(data), 0666)
	if e == nil {
		fmt.Println("data saved to file")
	} else {
		fmt.Println("data is not saved to file")
	}
	return e
}
func RetrieveFromFile(filename string) string {
	bs := Retrieve(filename)
	return string(bs)
}

func Retrieve(filename string) []byte {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bs
}
