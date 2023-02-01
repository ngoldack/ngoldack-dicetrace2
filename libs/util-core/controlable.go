package core

import "context"

type Controllable interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
