package repository

import (
	"main/database"
	"main/database/repository/auth"
)

type Repository interface {
	auth.AuthRepository
}

type repository struct {
	auth.AuthRepository
}

func NewRepository() Repository {
	db := database.DB
	return &repository{
		auth.NewAuthRepository(db),
	}
}
