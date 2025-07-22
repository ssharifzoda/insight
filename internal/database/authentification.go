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

func (a *AuthDb) UpdateAuthParams(userId int, sessionId string) error {
	var userAuth models.UserAuth
	userAuth.UserId = userId
	userAuth.SessionId = sessionId
	err := a.conn.Table("user_auth").Where("user_id", userId).Save(&userAuth).Error
	return err
}

func (a *AuthDb) GetAuthParamsByUserId(userId int) (*models.UserAuth, error) {
	var params *models.UserAuth
	err := a.conn.Table("user_auth").Where("user_id", userId).First(&params).Error
	return params, err
}

func (a *AuthDb) ChangeUserPassword(request *models.ChangePassword) error {
	tx := a.conn.Begin()
	err := tx.Table("users").Where("id", request.UserId).UpdateColumns(map[string]interface{}{
		"password":   request.NewPassword,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Table("user_auth").Where("user_id", request.UserId).UpdateColumns(map[string]interface{}{
		"temporary_pass": 0,
		"pass_reset_at":  time.Now().AddDate(0, 3, 0),
		"updated_at":     time.Now(),
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
