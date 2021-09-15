package main

import (
	"log"

	infrastructure "github.com/Ras96/clean-architecture-sample/3_infrastructure"
)

func main() {
	log.Println("Server started")

	e := infrastructure.InitRouting()
	e.Logger.Fatal(e.Start(":8080"))
}
