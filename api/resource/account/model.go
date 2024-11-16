package account

import (
	database "github.com/leondore/expense-tracker-api/database/gen"
)

type DTO = database.ListAccountsRow

type Form = database.CreateAccountParams
