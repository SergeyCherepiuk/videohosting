package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergeyCherepiuk/videohosting/pkg/http"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	e := http.NewRouter()
	e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
