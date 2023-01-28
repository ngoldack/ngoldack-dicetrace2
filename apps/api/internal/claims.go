package internal

import (
	"context"
)

// Claims contains custom data we want from the token.
type Claims struct {
}

// Validate errors out if `ShouldReject` is true.
func (c *Claims) Validate(ctx context.Context) error {
	return nil
}
