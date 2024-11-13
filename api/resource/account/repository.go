package account

import database "github.com/leondore/expense-tracker-api/database/gen"

type Repository struct {
	DB *database.Queries
}
