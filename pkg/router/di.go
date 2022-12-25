package router

import (
	"github.com/eminoz/go-api/pkg/api"
	"github.com/eminoz/go-api/pkg/repository"
	"github.com/eminoz/go-api/pkg/service"
)

type base struct{}

func (b base) UserDI() api.UserApi {
	r := repository.UserCollectionSetting()
	s := service.NewUserService(r)
	a := api.NewUserApi(s)
	return a
}
