package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leondore/expense-tracker-api/api/router"
	"github.com/leondore/expense-tracker-api/config"
)

//  @title          Expense Tracker API
//  @version        1.0
//  @description    Expense tracker RESTful API. Built with Golang.

//  @contact.name   Leon Dore
//  @contact.url    https://github.com/leondore

//  @license.name   MIT License
//  @license.url    https://github.com/leondore/expense-tracker-api/blob/master/LICENSE

// @host           localhost:8080
// @basePath       /v1
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
