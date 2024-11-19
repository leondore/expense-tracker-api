package account

import (
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

func NewDTO(row database.ListAccountsRow) *DTO {
	return &DTO{
		Id:            row.ID,
		Name:          row.Name,
		Balance:       types.Currency(row.Balance),
		Description:   &row.Description.String,
		AccountNumber: &row.AccountNumber.String,
		CategoryID:    int(row.CategoryID),
		Category:      row.Category,
		Institution:   row.Institution,
		Currency:      row.Currency,
	}
}
