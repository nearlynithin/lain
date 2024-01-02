package lain

import (
	"context"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"
	_ "github.com/lib/pq"
)

func TestService_CreatePost(t *testing.T) {
	svc := &Service{Queries: testQueries}
	ctx := context.Background()

	t.Run("empty_content", func(t *testing.T) {
		_, err := svc.CreatePost(ctx, CreatePostInput{Content: ""})
		assert.EqualError(t, err, "invalid post content")
	})

	t.Run("too_long_content", func(t *testing.T) {
		s := strings.Repeat("a", 1001)
		_, err := svc.CreatePost(ctx, CreatePostInput{Content: s})
		assert.EqualError(t, err, "invalid post content")
	})

	t.Run("unauthenticated", func(t *testing.T) {
		_, err := svc.CreatePost(ctx, CreatePostInput{Content: "Bhupendra Jogi"})
		assert.EqualError(t, err, "unauthenticated")
	})

	t.Run("ok", func(t *testing.T) {
		asUser := ContextWithUser(ctx, genUser(t))
		_, err := svc.CreatePost(asUser, CreatePostInput{Content: "Bhupendra Jogi"})
		assert.NoError(t, err)

	})
}

func genUser(t *testing.T) User {
	t.Helper()

	ctx := context.Background()
	userID := genID()
	email := "23c12.nithin@sjec.ac.in"
	username := "BhupendraJogi"
	createdAt, err := testQueries.CreateUser(ctx, CreateUserParams{
		UserID:   userID,
		Email:    email,
		Username: username,
	})
	assert.NoError(t, err)
	return User{
		ID:        userID,
		Email:     email,
		Username:  username,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
}

func TestService_Posts(t *testing.T) {
	svc := &Service{Queries: testQueries}
	ctx := context.Background()

	t.Run("ok", func(t *testing.T) {
		got, err := svc.Posts(ctx)
		assert.NoError(t, err)
		for _, p := range got {
			assert.NotZero(t, p)
		}
	})
}

func TestService_Post(t *testing.T) {
	svc := &Service{Queries: testQueries}
	ctx := context.Background()

	t.Run("invalid_id", func(t *testing.T) {
		_, err := svc.Post(ctx, "@nope@")
		assert.EqualError(t, err, "invalid post ID")
	})

	t.Run("not_found", func(t *testing.T) {
		_, err := svc.Post(ctx, genID())
		assert.EqualError(t, err, "post not found")
	})

}
