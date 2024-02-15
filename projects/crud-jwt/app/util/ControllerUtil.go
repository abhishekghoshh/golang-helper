package util

import (
	"crud/app/dto"
	"crud/app/error"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func AddRoute(router *mux.Router, route string, fn func(http.ResponseWriter, *http.Request), method string, routeName string) {
	router.
		HandleFunc(route, fn).
		Methods(method).
		Name(routeName)
}
func Write(w http.ResponseWriter, r *http.Request, data interface{}, statusCode int, appError *error.AppError) {
	w.Header().Add("Content-Type", "application/json")
	if nil != appError {
		w.WriteHeader(appError.Code)
		json.NewEncoder(w).Encode(appError.AsMessage())
	} else {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(data)
	}
}

func JsonDecode[T any](w http.ResponseWriter, r *http.Request, dataPtr *T) bool {
	err := json.NewDecoder(r.Body).Decode(dataPtr)
	if err != nil {
		Write(w, r, *dataPtr, http.StatusOK, error.NewUnexpectedError(err.Error()))
		return false
	}
	return true
}

func DecodeLoginRequest(loginRequest *dto.LoginRequest) *error.AppError {
	if decodedText, err := base64.StdEncoding.DecodeString(loginRequest.UserName); err != nil {
		return error.NewUnexpectedError(err.Error())
	} else {
		loginRequest.UserName = string(decodedText)
	}
	if decodedText, err := base64.StdEncoding.DecodeString(loginRequest.Password); err != nil {
		return error.NewUnexpectedError(err.Error())
	} else {
		loginRequest.Password = string(decodedText)
	}
	return nil
}
