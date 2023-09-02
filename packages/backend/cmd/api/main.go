package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/divyam-rai/personal-finance-tracker/packages/backend/internal/api"
)

func main() {
	router, err := api.New()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting HTTP server at port 8080\n")
	if err := http.ListenAndServe(":8080", router.Mux); err != nil {
		log.Fatal(err)
	}
}
