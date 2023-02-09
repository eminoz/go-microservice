package service

import (
	"testing"

	"github.com/eminoz/go-api/pkg/internal/mocks/brokermocks"
	"github.com/eminoz/go-api/pkg/internal/mocks/cachemocks"
	"github.com/eminoz/go-api/pkg/internal/mocks/encryptionmocks"
	"github.com/eminoz/go-api/pkg/internal/mocks/jwtmocks"
	"github.com/eminoz/go-api/pkg/internal/mocks/repomocks"
	"github.com/eminoz/go-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUserById(t *testing.T) {
	mockRepo := &repomocks.UserRepository{}
	brokermocks := &brokermocks.User{}
	cachemocks := &cachemocks.UserCache{}
	encryptionmocks := &encryptionmocks.Encryption{}
	jwtmocks := &jwtmocks.AuthJwt{}

	id := "1234hdhsj"
	user := model.User{Name: "emin", Email: "eminoz@gmail.com", Password: "123432d", Role: "user"}

	mockRepo.On("UpdateUserById", id, user).Return(true, "user updated")
	service := NewUserService(mockRepo, brokermocks, cachemocks, encryptionmocks, jwtmocks)
	a := service.UpdateUserById(id, &user)

	assert.Equal(t, a, "user updated")

}
