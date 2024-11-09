package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/leondore/expense-tracker-api/api/resource/account"
	"github.com/leondore/expense-tracker-api/api/resource/health"
)

func New() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/healthz", health.Read)

	r.Route("/v1", func(r chi.Router) {
		accountAPI := &account.API{}
		r.Get("/accounts", accountAPI.List)
		r.Post("/accounts", accountAPI.Create)
		r.Get("/accounts/{id}", accountAPI.Read)
		r.Put("/accounts/{id}", accountAPI.Update)
		r.Delete("/accounts/{id}", accountAPI.Delete)
	})

	return r
}
