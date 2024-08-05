package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	serv := http.FileServer(http.Dir("."))
	mux.Handle("/", serv)

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error starting server : %s\n", err)
	}
}
