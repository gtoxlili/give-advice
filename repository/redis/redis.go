package redis

import (
	"sync"

	"github.com/gtoxlili/advice-hub/constants"
	"github.com/redis/go-redis/v9"
)

var (
	redisOnce  sync.Once
	repository *Repository
)

type Repository struct {
	*redis.Client
}

func Default() *Repository {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     constants.RedisAddr,
			Password: constants.RedisPassword,
			DB:       0,
		})
		repository = &Repository{client}
	})
	return repository
}
