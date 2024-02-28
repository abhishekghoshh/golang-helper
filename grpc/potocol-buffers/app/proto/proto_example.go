package proto

import (
	"errors"
	"proto-example/app/proto/one_of"
	"proto-example/app/proto/sample"
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

func NewFullAddress() *sample.FullAddress {
	sampleGoogleLocation := "sample google location"
	return &sample.FullAddress{
		BuildingName:   "sample builing",
		StreetName:     "sample street",
		AreaName:       "sample area",
		CityName:       "sample city",
		DistrictName:   "sample sitrict",
		StateName:      "sample state",
		CountryName:    "sample country",
		PinCode:        123456,
		GoogleLocation: &sampleGoogleLocation,
		LandMarks: []string{
			"landmark 1",
			"landmark 2",
		},
		OtherInformations: map[string]string{
			"1": "sample 1",
			"2": "sample 2",
		},
		AddressType: sample.AddressType_ADDRESS_TYPE_URBAN,
	}
}
func EmptyFullAddress() *sample.FullAddress {
	return &sample.FullAddress{}
}

func NewSuccessResponse() *one_of.Response {
	successMessage := "wow!! this is a success"
	return &one_of.Response{
		Response: &one_of.Response_Success{
			Success: &one_of.Success{
				Mssg: []byte(successMessage),
			},
		},
		Code: 200,
		Headers: map[string]string{
			"mediaType": "application/json",
		},
	}
}

func NewErrorResponse() *one_of.Response {
	errorMessage := "noooo!! this is an error"
	return &one_of.Response{
		Response: &one_of.Response_Error{
			Error: &one_of.Error{
				Mssg: []byte(errorMessage),
				StackTrace: []string{
					"this is root error",
					"this is thown afterwards",
				},
			},
		},
		Code: 200,
		Headers: map[string]string{
			"mediaType": "application/json",
		},
	}
}
func DummyResponse() *one_of.Response {
	return &one_of.Response{}
}

func IsSuccess(response *one_of.Response) (bool, error) {
	switch response.Response.(type) {
	case *one_of.Response_Success:
		return true, nil
	case *one_of.Response_Error:
		return false, nil
	default:
		return false, errors.New("response is empty")
	}
}
