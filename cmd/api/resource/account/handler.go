package account

import "net/http"

type API struct{}

func (api *API) List(w http.ResponseWriter, r *http.Request) {}

func (api *API) Create(w http.ResponseWriter, r *http.Request) {}

func (api *API) Read(w http.ResponseWriter, r *http.Request) {}

func (api *API) Update(w http.ResponseWriter, r *http.Request) {}

func (api *API) Delete(w http.ResponseWriter, r *http.Request) {}
