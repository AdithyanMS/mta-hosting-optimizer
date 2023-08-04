package main

import (
	"fmt"
	"log"
	"net/http"

	"errors"

	"github.com/AdithyanMS/mta-hosting-optimizer/handler"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(errors.New("unable to load .env"))
	}

	port := 8081
	mux := http.NewServeMux()

	mux.HandleFunc("/inefficient_hosts", handler.InefficientHosts)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	fmt.Printf("server running at port %d\n", port)
	log.Fatal(server.ListenAndServe())
	fmt.Println("Server shutting down")
}
