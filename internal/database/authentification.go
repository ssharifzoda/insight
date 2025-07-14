package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"insight/pkg/consts"
	"time"
)

type AuthDb struct {
	conn *gorm.DB
}

func NewAuthDb(conn *gorm.DB) *AuthDb {
	return &AuthDb{conn: conn}
}

func (a *AuthDb) GetUserPermission(userId int) ([]int, error) {
	var permissions []int
	err := a.conn.Raw(consts.GetUserPermissionsByIdSQL, userId).Scan(&permissions).Error
	return permissions, err
}

func (a *AuthDb) UpdateRefreshToken(userId int, accessToken, refreshToken string) error {
	var userAuth *models.UserAuth
	userAuth.UserId = userId
	userAuth.RefreshToken = refreshToken
	userAuth.AccessToken = accessToken
	err := a.conn.Save(&userAuth).Error
	return err
}

func (a *AuthDb) GetTokenByUserId(userId int) (*models.UserAuth, error) {
	var params *models.UserAuth
	err := a.conn.Where("user_id", userId).First(&params).Error
	return params, err
}

func (a *AuthDb) ChangeUserPassword(request *models.ChangePassword) error {
	err := a.conn.Table("users").Where("id", request.UserId).UpdateColumns(map[string]interface{}{
		"password":   request.NewPassword,
		"updated_at": time.Now(),
	}).Error
	return err
}
