package account

import (
	"net/http"

	database "github.com/leondore/expense-tracker-api/database/gen"
)

type API struct {
	repository *Repository
}

func NewAPI(db *database.Queries) *API {
	return &API{
		repository: NewRepository(db),
	}
}

// List godoc
//
//	@summary        List accounts
//	@description    List accounts
//	@tags           accounts
//	@accept         json
//	@produce        json
//	@success        200 {array}     DTO
//	@failure        500 {object}    err.Error
//	@router         /accounts [get]
func (api *API) List(w http.ResponseWriter, r *http.Request) {
	accounts, err := api.repository.List(r.Context())
	if err != nil {
		// handle error
	}

	list := make([]DTO, 0, len(accounts))
	for _, account := range accounts {
		var dto DTO
		list = append(list, *dto.FromModel(account))
	}
}

// Create godoc
//
//	@summary        Create account
//	@description    Create account
//	@tags           accounts
//	@accept         json
//	@produce        json
//	@param          body    body    Form    true    "Account form"
//	@success        201
//	@failure        400 {object}    err.Error
//	@failure        422 {object}    err.Errors
//	@failure        500 {object}    err.Error
//	@router         /accounts [post]
func (api *API) Create(w http.ResponseWriter, r *http.Request) {}

// Read godoc
//
//	@summary        Read account
//	@description    Read account
//	@tags           accounts
//	@accept         json
//	@produce        json
//	@param          id	path        string  true    "Account ID"
//	@success        200 {object}    DTO
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /accounts/{id} [get]
func (api *API) Read(w http.ResponseWriter, r *http.Request) {}

// Update godoc
//
//	@summary        Update account
//	@description    Update account
//	@tags           accounts
//	@accept         json
//	@produce        json
//	@param          id      path    string  true    "Account ID"
//	@param          body    body    Form    true    "Account form"
//	@success        200
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        422 {object}    err.Errors
//	@failure        500 {object}    err.Error
//	@router         /accounts/{id} [put]
func (api *API) Update(w http.ResponseWriter, r *http.Request) {}

// Delete godoc
//
//	@summary        Delete account
//	@description    Delete accounts
//	@tags           accounts
//	@accept         json
//	@produce        json
//	@param          id  path    string  true    "Account ID"
//	@success        200
//	@failure        400 {object}    err.Error
//	@failure        404
//	@failure        500 {object}    err.Error
//	@router         /accounts/{id} [delete]
func (api *API) Delete(w http.ResponseWriter, r *http.Request) {}
