package lain

import (
	"context"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestService_CreateComment(t *testing.T) {
	svc := &Service{Queries: testQueries}
	ctx := context.Background()

	t.Run("invalid_post_id", func(t *testing.T) {
		_, err := svc.CreateComment(ctx, CreateCommentInput{PostID: "@nope@"})
		assert.EqualError(t, err, "invalid post ID")
	})

	t.Run("empty_content", func(t *testing.T) {
		_, err := svc.CreateComment(ctx, CreateCommentInput{
			PostID:  genID(),
			Content: "",
		})
		assert.EqualError(t, err, "invalid comment content")
	})

	t.Run("too_long_content", func(t *testing.T) {
		s := strings.Repeat("a", 1001)
		_, err := svc.CreateComment(ctx, CreateCommentInput{
			PostID:  genID(),
			Content: s,
		})
		assert.EqualError(t, err, "invalid comment content")
	})

	t.Run("unauthenticated", func(t *testing.T) {
		_, err := svc.CreateComment(ctx, CreateCommentInput{
			PostID:  genID(),
			Content: "Bhupendra Jogi",
		})
		assert.EqualError(t, err, "unauthenticated")
	})
}
