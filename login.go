//this is a different login file that is in the root of the directory

package lain

import (
	"context"
	"errors"

	"github.com/rs/xid"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrUsernameTaken = errors.New("username taken")
)

//making a login function that is a part of the service

type LoginInput struct {
	Email    string
	Username *string
}

func (svc *Service) Login(ctx context.Context, in LoginInput) (User, error) {
	var out User

	exists, err := svc.Queries.UserExistsByEmail(ctx, in.Email)

	if err != nil {
		return out, err
	}

	if exists {
		return svc.Queries.UserByEmail(ctx, in.Email)
	}

	if in.Username == nil {
		return out, ErrUserNotFound
	}

	exists, err = svc.Queries.UserExistsByUsername(ctx, *in.Username)
	if err != nil {
		return out, err
	}

	if exists {
		return out, ErrUsernameTaken
	}

	//finally, creating a user here
	userID := genID()
	createdAt, err := svc.Queries.CreateUser(ctx, CreateUserParams{
		UserID:   userID,
		Email:    in.Email,
		Username: *in.Username,
	})

	if err != nil {
		return out, err
	}

	return User{
		ID:        userID,
		Email:     in.Email,
		Username:  *in.Username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}, nil
}

// So now the login method is done, It can be used by the handler now
func genID() string {
	return xid.New().String()
}
