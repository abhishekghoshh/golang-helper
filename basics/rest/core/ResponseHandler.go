package core

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

func AddToResponse(w http.ResponseWriter, r *http.Request, data interface{}) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(data)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
