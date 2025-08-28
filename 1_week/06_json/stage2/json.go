package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Request struct {
	ID     int              `json:"id"`
	Name   string           `json:"name"`
	Cars   []Car            `json:"cars,omitempty"`//omitempty не будет ключа если пустое  
	Params map[string]Param `json:"params"`
}

type Car struct {
	Plate string `json:"plate"`
	Brand string `json:"brand"`
}

type Param struct {
	ValueID   int64  `json:"value_id"`
	ValueName string `json:"value_name"`
}

func main() {
	request := Request{
		ID: 123,
		Name: "Zubenko Mihail",
		Params: map[string]Param{
			"occupation": {
				ValueID: 900,
				ValueName: "mafia",
			},
		},
	}

	rawJSON, err := json.Marshal(request)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(rawJSON))
}