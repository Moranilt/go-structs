package request

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ParsedBodyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ParsedBody(r *http.Request) ParsedBodyStruct {
	var parsedBody ParsedBodyStruct
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {

	}
	fmt.Printf("Body %s", body)
	err = json.Unmarshal(body, &parsedBody)
	return parsedBody
}

func ToJSON(body ParsedBodyStruct) []byte {
	output, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	return output
}
