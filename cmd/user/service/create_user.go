package service

import (
	"context"
	"douyin/cmd/user/dao"
	"douyin/cmd/user/repository"
)

type CreateUserServiceImpl struct {
	ctx context.Context
}

func NewCreateUserServiceImpl(ctx context.Context) *CreateUserServiceImpl {
	return &CreateUserServiceImpl{ctx: ctx}
}

func (receiver *CreateUserServiceImpl) CreateUser(username string, password string) error {
	userDao := dao.GetUserDaoInstance()
	return userDao.CreateUser(&repository.User{
		Name:     username,
		Password: password,
	})
}
