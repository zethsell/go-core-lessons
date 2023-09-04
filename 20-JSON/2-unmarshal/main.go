package main

import (
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
	dog1JSON := `{"name":"Rx","breed":"Dalmata","age":3}`

	var dog1 dog

	fmt.Println(dog1)

	if err := json.Unmarshal([]byte(dog1JSON), &dog1); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog1)

	dog2JSON := `{"breed":"Poodle","name":"Toby"}`

	dog2 := make(map[string]string)

	fmt.Println(dog2)

	if err := json.Unmarshal([]byte(dog2JSON), &dog2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(dog2)

}
