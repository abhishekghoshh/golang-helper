package proto

import (
	with_pkg "proto-example/app/proto/segregated_with_package"
)

func AddressWithPackage() *with_pkg.Address {
	city := &with_pkg.City{
		Name:        "Berhampore",
		ZipCode:     "742101",
		CountryName: "India",
	}
	street := &with_pkg.Street{
		Name: "bagan Road",
		City: city,
	}
	building := &with_pkg.Building{
		Name:   "New Abason",
		Number: 100,
		Street: street,
	}
	return &with_pkg.Address{
		City:     city,
		Street:   street,
		Building: building,
	}
}

func EmptyAddressWithPackage() *with_pkg.Address {
	return &with_pkg.Address{}
}
