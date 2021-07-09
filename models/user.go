package models

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserModel struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone string `json:"phone"`
}

func User(r *http.Request) *UserModel {
	var user UserModel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	return &user
}

func (user *UserModel) ToJSON() []byte {
	output, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	return output
}
