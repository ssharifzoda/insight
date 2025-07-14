package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/utils"
)

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}

func (a *AuthService) GetUserPermission(userId int) ([]int, error) {
	return a.db.GetUserPermission(userId)
}
func (a *AuthService) UpdateRefreshToken(userId int, accessToken, refreshToken string) error {
	return a.db.UpdateRefreshToken(userId, accessToken, refreshToken)
}
func (a *AuthService) GetTokenByUserId(userId int) (*models.UserAuth, error) {
	return a.db.GetTokenByUserId(userId)
}
func (a *AuthService) ChangeUserPassword(request *models.ChangePassword) error {
	hash, err := utils.Hash(request.NewPassword)
	if err != nil {
		return err
	}
	request.NewPassword = hash
	return a.db.ChangeUserPassword(request)
}
