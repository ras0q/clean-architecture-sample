package main

import (
	"log"
	"regexp"
)

func main() {
	log.Println("Server started")

	// infrastructure.InitRouting()

	r := regexp.MustCompile(`a`)
	res := r.Split("banana", -1)
	log.Println(res, len(res))
}
