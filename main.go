package main

import "github.com/ultrazg/xyz/service"

func main() {
	err := service.Start()
	if err != nil {
		panic(err)
	}
}
