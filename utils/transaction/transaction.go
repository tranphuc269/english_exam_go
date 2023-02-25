package transaction

import "context"

type Transaction interface {
	Required(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
