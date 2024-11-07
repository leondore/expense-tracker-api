package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leondore/expense-tracker-api/cmd/api/router"
	"github.com/leondore/expense-tracker-api/config"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal(err.Error())
	}
	c := config.New()

	r := router.New()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	log.Println("Starting server " + s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}
