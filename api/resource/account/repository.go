package account

import (
	"context"

	"github.com/google/uuid"
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

func (r *Repository) List(ctx context.Context, userId string) ([]DTO, error) {
	uuid, err := uuid.Parse(userId)
	if err != nil {
		return nil, err
	}

	return r.db.ListAccounts(ctx, uuid)
}
