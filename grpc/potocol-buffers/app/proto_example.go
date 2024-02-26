package app

import (
	"proto-example/app/proto"
	"proto-example/app/proto/import_pkg"
	"proto-example/app/proto/imports"
)

func NewAccount() *proto.Account {
	return &proto.Account{
		Id:         1,
		Name:       "Abhishek Ghosh",
		IsVerified: true,
	}
}

func NewAddress1() *imports.Address {
	city := &imports.City{
		Name:        "Berhampore",
		ZipCode:     "742101",
		CountryName: "India",
	}
	street := &imports.Street{
		Name: "bagan Road",
		City: city,
	}
	building := &imports.Building{
		Name:   "New Abason",
		Number: 100,
		Street: street,
	}
	return &imports.Address{
		City:     city,
		Street:   street,
		Building: building,
	}
}

func NewAddress2() *import_pkg.Address {
	city := &import_pkg.City{
		Name:        "Berhampore",
		ZipCode:     "742101",
		CountryName: "India",
	}
	street := &import_pkg.Street{
		Name: "bagan Road",
		City: city,
	}
	building := &import_pkg.Building{
		Name:   "New Abason",
		Number: 100,
		Street: street,
	}
	return &import_pkg.Address{
		City:     city,
		Street:   street,
		Building: building,
	}
}
