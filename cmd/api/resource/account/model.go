package account

import "github.com/leondore/expense-tracker-api/utils/currency"

type DTO struct {
	Id            string            `json:"id"`
	Name          string            `json:"name"`
	Balance       currency.Currency `json:"balance"`
	Currency      string            `json:"currency"`
	Description   string            `json:"description"`
	Category      string            `json:"category"`
	AccountNumber string            `json:"account_number"`
	Institution   string            `json:"institution"`
}
