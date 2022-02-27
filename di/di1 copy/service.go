package main

import (
	"context"

	"github.com/ludwig125/ludwig125_gosample/di/di1/model"
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userService struct {
	ur UserRepository
}

func NewUserService(ur UserRepository) UserService {
	return &userService{ur}
}

func (us *userService) CreateUser(ctx context.Context, user *model.User) {
	us.ur.CreateUser(ctx, user)
}
