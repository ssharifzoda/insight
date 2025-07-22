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
func (a *AuthService) UpdateAuthParams(userId int, sessionId string) error {
	return a.db.UpdateAuthParams(userId, sessionId)
}
func (a *AuthService) GetAuthParamsByUserId(userId int) (*models.UserAuth, error) {
	return a.db.GetAuthParamsByUserId(userId)
}
func (a *AuthService) ChangeUserPassword(request *models.ChangePassword) error {
	hash, err := utils.Hash(request.NewPassword)
	if err != nil {
		return err
	}
	request.NewPassword = hash
	return a.db.ChangeUserPassword(request)
}
