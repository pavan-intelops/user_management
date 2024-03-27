package services

import (
	"github.com/pavan-intelops/user_management/user_management/pkg/rest/server/daos"
	"github.com/pavan-intelops/user_management/user_management/pkg/rest/server/models"
)

type UserService struct {
	userDao *daos.UserDao
}

func NewUserService() (*UserService, error) {
	userDao, err := daos.NewUserDao()
	if err != nil {
		return nil, err
	}
	return &UserService{
		userDao: userDao,
	}, nil
}
