package services

import (
	"context"
	"github.com/amirrezam75/kenopsiauser"
)

const authUserCtxKey = "authUser"

type ContextService struct{}

func (cs ContextService) WithUser(ctx context.Context, user kenopsiauser.User) context.Context {
	return context.WithValue(ctx, authUserCtxKey, user)
}

func (cs ContextService) GetUser(ctx context.Context) *kenopsiauser.User {
	user, ok := ctx.Value(authUserCtxKey).(*kenopsiauser.User)

	if !ok {
		return nil
	}

	return user
}
