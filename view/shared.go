package view

import (
	"accent-ui/types"
	"context"
)

// authenticatedUser returns the authenticated user from the context
//
func AuthenticatedUser(ctx context.Context) types.AuthenticatedUser {
	user, ok := ctx.Value(types.UserContextKey).(types.AuthenticatedUser)
	if !ok {
		return types.AuthenticatedUser{}
	}
	return user
}