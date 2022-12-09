package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	infrastructure "github.com/ras0q/clean-architecture-sample/3_infrastructure"
	alt "github.com/ras0q/clean-architecture-sample/3_infrastructure_alt"
)

func main() {
	log.Println("Server started")

	port, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		port = 8080
	}

	// 使用するフレームワークを変更する
	if os.Getenv("APP_INFRA") == "alt" {
		// gin
		r := alt.InitRouting()
		r.Run(fmt.Sprintf(":%d", port))
	} else {
		// echo
		e := infrastructure.InitRouting()
		e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
	}
}
