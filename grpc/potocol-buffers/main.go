package main

import (
	"fmt"
	app "proto-example/app/proto"
	"reflect"

	fh "proto-example/app/file_handling"

	"google.golang.org/protobuf/proto"
)

func main() {
	address := app.NewFullAddress()
	// fmt.Println(address)
	DoFile(address)
	// DoJson(address)
	// DoOneOfTest()
}

func DoOneOfTest() {
	isSuccess, err := app.IsSuccess(app.NewSuccessResponse())
	fmt.Println("Is this a success reponse :", isSuccess, "has any error :", err)

	isSuccess, err = app.IsSuccess(app.NewErrorResponse())
	fmt.Println("Is this a success reponse :", isSuccess, "has any error :", err)

	isSuccess, err = app.IsSuccess(app.DummyResponse())
	fmt.Println("Is this a success reponse :", isSuccess, "has any error :", err)
}

func DoFile(message proto.Message) {
	path := "./files/simple.bin"
	fh.WriteToFile(path, message)

	newMessageType := app.EmptyFullAddress()
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
