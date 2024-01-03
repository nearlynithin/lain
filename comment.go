package lain

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nicolasparada/go-errs"
)

const (
	ErrInvalidCommentContent = errs.InvalidArgumentError("invalid comment content")
)

type CreateCommentInput struct {
	PostID  string
	Content string
}

func (in *CreateCommentInput) Prepare() {
	in.Content = strings.TrimSpace(in.Content)
	in.Content = strings.ReplaceAll(in.Content, "\n\n", "\n")
	in.Content = strings.ReplaceAll(in.Content, "  ", " ")
}

func (in CreateCommentInput) Validate() error {
	//same as validating the post content, but here we are validating the post id too
	if !isID(in.PostID) {
		return ErrInvalidPostID
	}
	if in.Content == "" || utf8.RuneCountInString(in.Content) > 1000 {
		return ErrInvalidCommentContent
	}
	return nil
}

type CreateCommentOutput struct {
	ID        string
	CreatedAt time.Time
}

// a method
func (svc *Service) CreateComment(ctx context.Context, in CreateCommentInput) (CreateCommentOutput, error) {
	var out CreateCommentOutput

	in.Prepare()
	if err := in.Validate(); err != nil {
		return out, err
	}
	usr, ok := UserFromContext(ctx)
	if !ok {
		return out, errs.Unauthenticated
	}

	//finally, creating a comment
	commentID := genID()
	createdAt, err := svc.Queries.CreateComment(ctx, CreateCommentParams{
		CommentID: commentID,
		PostID:    in.PostID,
		UserID:    usr.ID,
		Content:   in.Content,
	})
	if err != nil {
		return out, err
	}
	out.ID = commentID
	out.CreatedAt = createdAt

	return out, nil
}

// a method to list a comments

func (svc *Service) Comments(ctx context.Context, postID string) ([]CommentsRow, error) {
	if !isID(postID) {
		return nil, ErrInvalidPostID
	}
	return svc.Queries.Comments(ctx, postID)
}
