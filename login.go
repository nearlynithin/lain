//this is a different login file that is in the root of the directory

package lain

import (
	"context"
	"regexp"
	"strings"

	_ "github.com/lib/pq"
	"github.com/nicolasparada/go-errs"
	"github.com/rs/xid"
)

const (
	ErrUserNotFound    = errs.NotFoundError("user not found")
	ErrUsernameTaken   = errs.ConflictError("username taken")
	ErrInvalidEmail    = errs.InvalidArgumentError("invalid email")
	ErrInvalidUsername = errs.InvalidArgumentError("invalid username")
)

var (
	reEmail    = regexp.MustCompile(`^\d{2}[a-zA-Z]\d{2}\.\w+@sjec\.ac\.in$`)
	reUsername = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_-]{0,17}`)
)

//making a login function that is a part of the service

type LoginInput struct {
	Email    string
	Username *string
}

func (in *LoginInput) Prepare() {
	in.Email = strings.ToLower(in.Email)
}

// using regex for email validation
func (in LoginInput) Validate() error {
	if !isEmail(in.Email) {
		return ErrInvalidEmail
	}

	if in.Username != nil && !isUsername(*in.Username) {
		return ErrInvalidUsername
	}

	return nil
}

func (svc *Service) Login(ctx context.Context, in LoginInput) (User, error) {
	var out User

	in.Prepare()

	if err := in.Validate(); err != nil {
		return out, err
	}

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

func isEmail(s string) bool {
	return reEmail.MatchString(s)
}
func isUsername(s string) bool {
	return reUsername.MatchString(s)
}

// So now the login method is done, It can be used by the handler now
func genID() string {
	return xid.New().String()
}
