package database

import (
	"errors"
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

func (u *UserDb) AddNewUser(params *models.User) (*models.User, error) {
	tx := u.conn.Begin()
	err := tx.Table("users").Create(&params).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	passReset := time.Now().AddDate(0, 3, 0)
	auth := &models.UserAuth{
		UserId:        params.Id,
		PassResetAt:   &passReset,
		TemporaryPass: 1,
	}
	err = tx.Table("user_auth").Create(&auth).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return params, nil
}

func (u *UserDb) UpdateUserParams(params *models.User) error {
	return u.conn.Table("users").Updates(&params).Error
}

func (u *UserDb) GetAllUsers(limit, offset int) ([]*models.User, int, error) {
	var (
		users []*models.User
		count int64
	)
	tx := u.conn.Table("users").Where("role_id not in (4,5) and active = 1").Limit(limit).Offset(offset).Find(&users)
	tx.Count(&count)
	return users, int(count), tx.Error
}

func (u *UserDb) GetUserById(userId int) (*models.UserInfo, error) {
	var (
		user *models.UserInfo
	)
	err := u.conn.Table("users").Where("active = 1 and id = ?", userId).First(&user).Error
	switch {
	case user.ShopId != 0:
		var shop *models.Shop
		err = u.conn.Table("shops").Where("id = ? and status = 1", user.ShopId).First(&shop).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		user.Shop = shop
		var salePoint *models.SalePoint
		err = u.conn.Table("sale_points").Where("shop_id = ? and status = 1", user.ShopId).First(&salePoint).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		user.SalePoint = salePoint
	case user.SupplierId != 0:
		var supplier *models.Supplier
		err = u.conn.Table("suppliers").Where("id = ? and status = 1", user.SupplierId).First(&supplier).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		}
		user.Supplier = supplier
	}
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
