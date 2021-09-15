package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	infrastructure "github.com/Ras96/clean-architecture-sample/3_infrastructure"
)

func main() {
	log.Println("Server started")

	e := infrastructure.InitRouting()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8080
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
