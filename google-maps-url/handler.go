package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		input = body
	}

	location := &Location{}
	if err := json.Unmarshal(input, location); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("https://maps.google.com/?q=%f,%f", location.Latitude, location.Longitude)))
}
