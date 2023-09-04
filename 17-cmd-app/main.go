package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("start")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
