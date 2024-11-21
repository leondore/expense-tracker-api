package account

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/leondore/expense-tracker-api/api/resource/common/types"
	database "github.com/leondore/expense-tracker-api/database/gen"
)

type DTO struct {
	Id            uuid.UUID                       `json:"id"`
	Name          string                          `json:"name"`
	Balance       types.Currency                  `json:"balance"`
	Description   *string                         `json:"description"`
	AccountNumber *string                         `json:"account_number"`
	CategoryID    int                             `json:"category_id"`
	Category      database.AccountCategoryOptions `json:"category"`
	Institution   string                          `json:"institution"`
	Currency      string                          `json:"currency"`
}

type Form struct {
	Name          string    `json:"name"`
	Description   *string   `json:"description"`
	AccountNumber *string   `json:"account_number"`
	CategoryID    int       `json:"category_id"`
	InstitutionID *int      `json:"institution"`
	UserID        uuid.UUID `json:"user_id"`
	CurrencyID    int       `json:"currency"`
}

func (d *DTO) FromModel(row database.ListAccountsRow) *DTO {
	return &DTO{
		Id:            row.ID,
		Name:          row.Name,
		Balance:       types.CurrencyFromString(row.Balance),
		Description:   types.NullableStringtoPtr(row.Description),
		AccountNumber: types.NullableStringtoPtr(row.AccountNumber),
		CategoryID:    int(row.CategoryID),
		Category:      row.Category,
		Institution:   row.Institution,
		Currency:      row.Currency,
	}
}

type CreateAccountParams struct {
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	AccountNumber sql.NullString `json:"account_number"`
	CategoryID    int16          `json:"category_id"`
	InstitutionID sql.NullInt16  `json:"institution_id"`
	UserID        uuid.UUID      `json:"user_id"`
	CurrencyID    int16          `json:"currency_id"`
}

func (f *Form) ToModel() *database.CreateAccountParams {
	return &database.CreateAccountParams{
		Name:          f.Name,
		Description:   types.PtrToNullableString(f.Description),
		AccountNumber: types.PtrToNullableString(f.AccountNumber),
		CategoryID:    int16(f.CategoryID),
		InstitutionID: types.PtrToNullableInt16(f.InstitutionID),
		UserID:        f.UserID,
		CurrencyID:    int16(f.CurrencyID),
	}
}
