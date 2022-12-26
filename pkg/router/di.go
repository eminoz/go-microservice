package router

import (
	"github.com/eminoz/go-api/pkg/api"
	"github.com/eminoz/go-api/pkg/broker"
	"github.com/eminoz/go-api/pkg/cache"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/service"
)

type base struct{}

var redis = cache.InitRedis() //connect redis
var userbroker = broker.NewUserProducer()

func (b base) UserDI() api.UserApi {
	userCache := cache.NewUserCache(redis) //create user redis client
	r := repository.UserCollectionSetting()
	s := service.NewUserService(r, userbroker, userCache)
	a := api.NewUserApi(s)
	return a
}
