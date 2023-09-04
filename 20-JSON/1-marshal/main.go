package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type dog struct {
	Name  string `json:"name"`
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	dog1 := dog{"Rx", "Dalmata", 3}
	fmt.Println(dog1)

	dog1JSON, err := json.Marshal(dog1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog1JSON)

	fmt.Println(bytes.NewBuffer(dog1JSON))

	dog2 := map[string]string{
		"name":  "Toby",
		"breed": "Poodle",
	}

	dog2JSON, err := json.Marshal(dog2)
	if err != nil {
		log.Fatal(err)
	}
	println(dog2JSON)

	fmt.Println(bytes.NewBuffer(dog2JSON))
}
