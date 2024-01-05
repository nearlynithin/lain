//this is a different login file that is in the root of the directory

package lain

import (
	"context"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestService_Login(t *testing.T) {
	svc := &Service{Queries: testQueries}
	ctx := context.Background()

	t.Run("invalid_email", func(t *testing.T) {
		_, err := svc.Login(ctx, LoginInput{Email: "nope"})
		assert.EqualError(t, err, "invalid email")
	})

	t.Run("invalid_username", func(t *testing.T) {
		_, err := svc.Login(ctx, LoginInput{Email: "23c12.nithin@sjec.ac.in", Username: ptr("heloo")})
		assert.EqualError(t, err, "invalid username")
	})

	t.Run("user_not_found", func(t *testing.T) {
		_, err := svc.Login(ctx, LoginInput{Email: "23c12.nithin@sjec.ac.in"})
		assert.EqualError(t, err, "user not found")
	})

	// t.Run("ok", func(t *testing.T) {
	// 	_, err := svc.Login(ctx, LoginInput{Email: "23c12.nithin@sjec.ac.in"})
	// 	assert.NoError(t, err)
	// })

}

func ptr[T any](v T) *T {
	return &v
}
