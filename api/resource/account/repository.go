package account

import (
	"context"

	e "github.com/leondore/expense-tracker-api/api/resource/common/err"
	"github.com/leondore/expense-tracker-api/api/resource/common/utils"
	database "github.com/leondore/expense-tracker-api/database/gen"
)

type Repository struct {
	db *database.Queries
}

func NewRepository(db *database.Queries) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(ctx context.Context) ([]database.ListAccountsRow, error) {
	uuid, err := utils.UserIDFromContext(ctx)
	if err != nil {
		return nil, e.NewErrUserID(err)
	}

	return r.db.ListAccounts(ctx, uuid)
}
