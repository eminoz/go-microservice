package cache

import (
	"context"
	"encoding/json"

	"github.com/eminoz/go-api/pkg/model"
	"github.com/go-redis/redis/v8"
)

type UserCache interface {
	SaveUserByID(id string, user model.UserDto)
	GetUSerById(id string) model.UserDto
	DeleteUserById(id string)
}
type userCache struct {
	Redis *redis.Client
}

func NewUserCache(r *redis.Client) UserCache {
	return &userCache{
		Redis: r,
	}
}

// func (u userCache) SaveUserEmailByID(id string, email string) {
// 	ctx := context.Background()
// 	err := u.Redis.Set(ctx, id, email, 0).Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
func (u userCache) SaveUserByID(id string, user model.UserDto) {
	ctx := context.Background()
	marshaledUser, _ := json.Marshal(user)
	u.Redis.HSet(ctx, "userInfos", id, marshaledUser)

}
func (u userCache) GetUSerById(id string) model.UserDto {
	ctx := context.Background()
	responseUser := u.Redis.HGet(ctx, "userInfos", id)
	var user model.UserDto
	json.Unmarshal([]byte(responseUser.Val()), &user)
	return user
}
func (u userCache) DeleteUserById(id string) {
	ctx := context.Background()
	u.Redis.HDel(ctx, "userInfos", id)
}
