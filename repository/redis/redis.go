package redis

import (
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisOnce  sync.Once
	repository *Repository
	Addr       = ""
	Password   = ""
)

type Repository struct {
	*redis.Client
}

func Default() *Repository {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     Addr,
			Password: Password,
			DB:       0,
		})
		repository = &Repository{client}
	})
	return repository
}
