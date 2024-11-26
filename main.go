package main

import (
	"github.com/ultrazg/xyz/service"
	"log"
)

func main() {
	err := service.Start()
	if err != nil {
		log.Fatal(err)
	}
}
