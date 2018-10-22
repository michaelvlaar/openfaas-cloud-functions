package function

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

type Request struct {
	Point   *Location   `json:"point"`
	Polygon []*Location `json:"polygon"`
}

type Response struct {
	PointInPolygon bool `json:"pointInPolygon"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		input = body
	}

	request := &Request{}

	if err := json.Unmarshal(input, request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pointInPolygon := false
	for i, j := 0, len(request.Polygon)-1; i < len(request.Polygon); i = i + 1 {
		if (((request.Polygon[i].Latitude <= request.Point.Latitude) && (request.Point.Latitude < request.Polygon[j].Latitude)) || ((request.Polygon[j].Latitude <= request.Point.Latitude) && (request.Point.Latitude < request.Polygon[i].Latitude))) && (request.Point.Longitude < (request.Polygon[j].Longitude-request.Polygon[i].Longitude)*(request.Point.Latitude-request.Polygon[i].Latitude)/(request.Polygon[j].Latitude-request.Polygon[i].Latitude)+request.Polygon[i].Longitude) {
			pointInPolygon = !pointInPolygon
		}
		j = i
	}

	result, err := json.Marshal(&Response{
		PointInPolygon: pointInPolygon,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
