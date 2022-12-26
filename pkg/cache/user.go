package cache

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type UserCache interface {
	SaveUserEmailByID(id string, email string)
}
type userCache struct {
	Redis *redis.Client
}

func NewUserCache(r *redis.Client) UserCache {
	return &userCache{
		Redis: r,
	}
}
func (u userCache) SaveUserEmailByID(id string, email string) {
	ctx := context.Background()
	err := u.Redis.Set(ctx, id, email, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}
