package persistence

import "context"

type UserRepository interface {
	Get(ctx context.Context)
}
