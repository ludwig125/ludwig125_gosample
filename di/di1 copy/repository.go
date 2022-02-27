package main

import (
	"context"
	"database/sql"

	"github.com/ludwig125/ludwig125_gosample/di/di1/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User)
}

type userRepository struct {
	db sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{*db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *model.User) {
	ur.db.Query("INSERT INTO テーブル名（列名1,列名2,……）")
}
