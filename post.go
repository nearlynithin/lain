package lain

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/nicolasparada/go-errs"
)

type CreatePostInput struct {
	Content string
}

func (in *CreatePostInput) Prepare() {
	//This will remove the empty spaces from the start and end of the whatever content
	in.Content = strings.TrimSpace(in.Content)
	//replacing all the double line breaks with single line breaks
	in.Content = strings.ReplaceAll(in.Content, "\n\n", "\n")
	//similarly double space with single spaces
	in.Content = strings.ReplaceAll(in.Content, "  ", " ")
}

func (in CreatePostInput) Validate() error {
	//making sure that the content is not empty and it's not less than maybe a 1000 characters.
	if in.Content == "" || utf8.RuneCountInString(in.Content) > 1000 {
		return errs.InvalidArgumentError("invalid post content")
	}
	return nil
}

type CreatePostOutput struct {
	ID       string
	CreateAt time.Time
}

func (svc *Service) CreatePost(ctx context.Context, in CreatePostInput) (CreatePostOutput, error) {
	var out CreatePostOutput

	in.Prepare()
	if err := in.Validate(); err != nil {
		return out, err
	}

	usr, ok := UserFromContext(ctx)
	if !ok {
		return out, errs.Unauthenticated
	}

	postID := genID()
	createdAt, err := svc.Queries.CreatePost(ctx, CreatePostParams{
		PostID:  postID,
		UserID:  usr.ID,
		Content: in.Content,
	})
	if err != nil {
		return out, err
	}

	out.ID = postID
	out.CreateAt = createdAt

	return out, nil
}
