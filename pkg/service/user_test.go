package service

import (
	"testing"

	"github.com/eminoz/go-api/pkg/internal/mocks/brokermocks"
	"github.com/eminoz/go-api/pkg/internal/mocks/repomocks"
	"github.com/eminoz/go-api/pkg/model"
)

func TestUpdateUserById(t *testing.T) {
	mockRepo := &repomocks.UserRepository{}
	brokermocks := &brokermocks.User{}
	id := "1234hdhsj"
	user := model.User{Name: "emn"}
	mockRepo.On("UpdateUserById", id, user).Return()

}
