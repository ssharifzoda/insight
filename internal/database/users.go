package database

import (
	"gorm.io/gorm"
	"insight/internal/models"
	"time"
)

type UserDb struct {
	conn *gorm.DB
}

func NewUserDb(conn *gorm.DB) *UserDb {
	return &UserDb{conn: conn}
}

func (u *UserDb) AddNewUser(params *models.User) error {
	return u.conn.Table("users").Create(&params).Error
}

func (u *UserDb) UpdateUserParams(params *models.User) error {
	return u.conn.Table("users").Updates(&params).Error
}

func (u *UserDb) GetAllUsers(limit, offset int) ([]*models.User, error) {
	var users []*models.User
	err := u.conn.Table("users").Where("active", true).Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (u *UserDb) GetUserById(userId int) (*models.User, error) {
	var user *models.User
	err := u.conn.Table("users").Where("active = 1 and id = ?", userId).First(&user).Error
	return user, err
}

func (u *UserDb) GetUserByPhone(phone string) (*models.User, error) {
	var user *models.User
	err := u.conn.Table("users").Where("phone", phone).First(&user).Error
	return user, err
}

func (u *UserDb) DeleteUser(userId int) error {
	return u.conn.Table("users").Where("id", userId).UpdateColumns(map[string]interface{}{
		"active":     0,
		"deleted_at": time.Now(),
	}).Error
}
