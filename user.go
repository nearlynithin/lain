package lain

import (
	"context"
	"database/sql"
	"errors"
)

func (svc *Service) UserByUsername(ctx context.Context, username string) (UserByUsernameRow, error) {
	var out UserByUsernameRow
	if !isUsername(username) {
		return out, ErrInvalidUsername
	}
	usr, _ := UserFromContext(ctx)

	out, err := svc.Queries.UserByUsername(ctx, UserByUsernameParams{
		FollowerID: usr.ID,
		Username:   username,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return out, ErrUserNotFound
	}
	return out, err
}
