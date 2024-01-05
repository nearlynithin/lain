package lain

import (
	"context"

	"github.com/nicolasparada/go-errs"
)

const (
	ErrCannotFollowSelf = errs.InvalidArgumentError("cannot follow self")
)

// a service to follow the user
func (svc *Service) FollowUser(ctx context.Context, followUserID string) error {
	if !isID(followUserID) {
		return ErrInvalidUserID
	}

	usr, ok := UserFromContext(ctx)
	if !ok {
		return errs.Unauthenticated
	}

	if usr.ID == followUserID {
		return ErrCannotFollowSelf
	}

	//making sure that the follow doesn't already exists
	exists, err := svc.Queries.UserFollowExists(ctx, UserFollowExistsParams{
		FollowerID: usr.ID,
		FollowedID: followUserID,
	})

	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	exists, err = svc.Queries.UserExists(ctx, followUserID)
	if err != nil {
		return err
	}
	if !exists {
		return ErrUserNotFound
	}

	//finally establishing the user followers relationship
	_, err = svc.Queries.CreateUserFollow(ctx, CreateUserFollowParams{
		FollowerID: usr.ID,
		FollowedID: followUserID,
	})

	if err != nil {
		return err
	}
	/*and now increasing the following count of the Authenticated user (usr.ID)*/
	_, err = svc.Queries.UpdateUser(ctx, UpdateUserParams{
		UserID:                   usr.ID,
		IncreaseFollowingCountBy: 1,
	})

	if err != nil {
		return err
	}

	/*Similarly increasing followers count of target user (followUserID)*/
	_, err = svc.Queries.UpdateUser(ctx, UpdateUserParams{
		UserID:                   followUserID,
		IncreaseFollowersCountBy: 1,
	})

	if err != nil {
		return err
	}

	return nil
}
