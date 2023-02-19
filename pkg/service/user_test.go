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
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func UserEnvs() (*repomocks.UserRepository, *brokermocks.User, *cachemocks.UserCache, *encryptionmocks.Encryption, *jwtmocks.AuthJwt) {
	mockRepo := &repomocks.UserRepository{}
	brokermocks := &brokermocks.User{}
	cachemocks := &cachemocks.UserCache{}
	encryptionmocks := &encryptionmocks.Encryption{}
	jwtmocks := &jwtmocks.AuthJwt{}
	return mockRepo, brokermocks, cachemocks, encryptionmocks, jwtmocks
}
func TestCreateUser(t *testing.T) {
	mockRepo, brokermocks, cachemocks, encryptionmocks, jwtmocks := UserEnvs()
	id := primitive.NewObjectID()
	user := model.User{ID: id, Name: "emin", Email: "eminoz@gmail.com", Password: "1234567", Role: "user"}
	u := model.UserDto{ID: id, Name: user.Name, Email: user.Email}

	brokermocks.On("CreatedUser", user)

	encryptionmocks.On("GenerateHashPassword", user.Password).Return("1234567", nil)
	cachemocks.On("SaveUserByID", id.Hex(), u)

	jwtmocks.On("CreateToken", user.Email, user.Password).Return(model.Token{ID: primitive.NewObjectID(), Email: user.Email, Role: "user", TokenString: "63e66aa02c0c4925438a4f66"}, nil)

	mockRepo.On("CreateUser", &user).Return(u)

	s := NewUserService(mockRepo, brokermocks, cachemocks, encryptionmocks, jwtmocks)
	dto, _ := s.CreateUser(&user)

	assert.Equal(t, dto.Email, user.Email)
}
func TestSignIn(t *testing.T) {
	a, b, c, d, e := UserEnvs()

	id := primitive.NewObjectID()
	user := model.User{ID: id, Name: "emin", Email: "eminoz1@gmail.com", Password: "1234567", Role: "user"}
	createUserResponse := model.Token{ID: id, Email: user.Email, Role: "user", TokenString: "63e66aa02c0c4925438a4f66"}
	dto := model.UserDto{ID: id, Name: user.Name, Email: user.Email}
	e.On("CreateToken", user.Email, user.Password).Return(createUserResponse, nil)

	a.On("GetUserByEmail", user.Email).Return(dto)
	u := NewUserService(a, b, c, d, e)
	authmodel := model.Authentication{Email: user.Email, Password: user.Password}
	responseUser, _ := u.SignIn(&authmodel)
	assert.Equal(t, responseUser.Email, createUserResponse.Email)

	a2, b2, c2, d2, e2 := UserEnvs()
	a2.On("GetUserByEmail", user.Email).Return(model.UserDto{}, nil)
	u2 := NewUserService(a2, b2, c2, d2, e2)
	responseUser2, _ := u2.SignIn(&authmodel)
	assert.Equal(t, responseUser2.UserDto.Email, "")
}
func TestGetUser(t *testing.T) {
	a, b, c, d, e := UserEnvs()
	id := primitive.NewObjectID()
	user := model.User{ID: id, Name: "emin", Email: "eminoz@gmail.com", Password: "1234567", Role: "user"}
	userDto := model.UserDto{ID: id, Name: user.Name, Email: user.Email}

	c.On("GetUSerById", id.Hex()).Return(userDto)
	a.On("GetUserByID", id.Hex()).Return(userDto)
	c.On("SaveUserByID", id.Hex(), userDto)
	u := NewUserService(a, b, c, d, e)
	userRes, _ := u.GetUser(id.Hex())
	assert.Equal(t, userRes.Email, user.Email)

	id2 := primitive.NewObjectID()
	user2 := model.User{ID: id2, Name: "emin", Email: "eminoz@gmail.com", Password: "1234567", Role: "user"}
	userDto2 := model.UserDto{ID: id2, Name: user2.Name}
	c.On("GetUSerById", id.Hex()).Return(userDto2)
	a.On("GetUserByID", id.Hex()).Return(userDto2)
	c.On("SaveUserByID", id.Hex(), userDto2)
	u2 := NewUserService(a, b, c, d, e)
	userRes2, _ := u2.GetUser(id.Hex())
	assert.Equal(t, userRes2.Email, user.Email, "got user")

}
