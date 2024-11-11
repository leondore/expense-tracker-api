package account

import (
	"github.com/google/uuid"
	"github.com/leondore/expense-tracker-api/utils/currency"
)

type DTO struct {
	Id            uuid.UUID         `json:"id"`
	Name          string            `json:"name"`
	Balance       currency.Currency `json:"balance"`
	Currency      string            `json:"currency"`
	Description   string            `json:"description"`
	Category      string            `json:"category"`
	AccountNumber string            `json:"account_number"`
	Institution   string            `json:"institution"`
}

type Form struct {
	Name          string    `json:"name"`
	Currency      int       `json:"currency"`
	Description   string    `json:"description"`
	CategoryId    int       `json:"category_id"`
	AccountNumber string    `json:"account_number"`
	InstitutionId int       `json:"institution"`
	UserId        uuid.UUID `json:"user_id"`
}
