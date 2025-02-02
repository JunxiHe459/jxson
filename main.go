package main

import (
	"json/json/encode"
)

type Example struct {
	A string `json:"A"`
	B int    `json:",string"`
	C []byte `json:",omitempty"`
	D struct {
		e string
		F string `json:"f"`
	}
}

func main() {
	data, err := encode.Marshal(1)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
