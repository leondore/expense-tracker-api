package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/leondore/expense-tracker-api/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err.Error())
	}
	c := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      mux,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}
