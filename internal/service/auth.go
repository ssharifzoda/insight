package service

import "insight/internal/database"

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}
