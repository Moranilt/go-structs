package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Coordinates struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type LocationsModel struct {
	Coordinates Coordinates `json:"coordinates"`
	PlaceName   string      `json:"place_name"`
}

func Location(r *http.Request) *LocationsModel {
	var location LocationsModel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &location)
	if err != nil {
		panic(err)
	}
	return &location
}

func (location *LocationsModel) ToJSON() []byte {
	output, err := json.Marshal(location)
	if err != nil {
		panic(err)
	}
	return output
}
