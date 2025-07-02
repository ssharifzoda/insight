package service

import (
	"errors"
	"green/internal/database"
	"green/internal/models"
	"green/pkg/consts"
	"green/pkg/utils"
)

type AuthService struct {
	db database.Authorization
}

func NewAuthService(db database.Authorization) *AuthService {
	return &AuthService{db: db}
}
func (a *AuthService) GetUserByUsername(username string) (*models.User, error) {
	user, err := a.db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New(consts.UserNotFound)
	}
	deHashPass, err := utils.DeHash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = deHashPass
	return user, nil
}

func (a *AuthService) UpdateRefreshToken(username, accessToken, refreshToken string) error {
	var auth models.UserAuth
	auth.Username = username
	auth.AccessToken = accessToken
	auth.RefreshToken = refreshToken
	return a.db.UpdateRefreshToken(&auth)
}

func (a *AuthService) ChangeUserPassword(request *models.ChangePassword) error {
	hash, err := utils.Hash(request.NewPassword)
	if err != nil {
		return err
	}
	request.NewPassword = hash
	return a.db.ChangeUserPassword(request)
}

func (a *AuthService) GetTokensByUsername(username string) (string, string, error) {
	return a.db.GetTokensByUsername(username)
}

func (a *AuthService) GetUserPermissionsByUsername(username string) ([]string, error) {
	return a.db.GetUserPermissionsByUsername(username)
}
