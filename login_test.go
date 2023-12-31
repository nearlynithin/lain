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

	t.Run("invalid email", func(t *testing.T) {
		_, err := svc.Login(ctx, LoginInput{Email: "nope"})
		assert.EqualError(t, err, "invalid email")
	})
}
