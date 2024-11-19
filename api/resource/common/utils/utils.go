package utils

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

func UserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return uuid.UUID{}, errors.New("user id not found in context, or not a string")
	}

	return uuid.Parse(userId)
}
