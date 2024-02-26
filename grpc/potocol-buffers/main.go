package main

import (
	"fmt"
	app "proto-example/app/proto"
	"reflect"

	fh "proto-example/app/file_handling"

	"google.golang.org/protobuf/proto"
)

func main() {
	address := app.AddressWithPackage()
	// fmt.Println(address)
	// DoFile(address)
	DoJson(address)
}

func DoFile(message proto.Message) {
	path := "./files/simple.bin"
	fh.WriteToFile(path, message)

	newMessageType := app.EmptyAddressWithPackage()
	fmt.Println("before reading :", newMessageType)

	fh.ReadFromFile(path, newMessageType)
	fmt.Println("after reading the content:", newMessageType)
}

func DoJson(message proto.Message) {
	jsonString := fh.ToJSON(app.AddressWithPackage())
	fmt.Println("json of the message is", jsonString)

	newMessageType := app.EmptyAddressWithPackage()
	typ := reflect.TypeOf(newMessageType)
	fmt.Println("type of the message is", typ)

	// newMessage := reflect.New(typ).Interface().(proto.Message)

	fmt.Println("before writing from json string the message is", newMessageType)

	fh.FromJSON(jsonString, newMessageType)
	fmt.Println("before writing from json string the message is", newMessageType)
}
