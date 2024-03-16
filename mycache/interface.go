package mycache

import "context"

type Cache interface {
	Set(ctx context.Context, key string, value interface{}, cost int64) error
	Get(ctx context.Context, key string) (interface{}, error)
}
