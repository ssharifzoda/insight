package service

import (
	"insight/internal/database"
	"insight/internal/models"
	"insight/pkg/utils"
)

type UserService struct {
	db database.Users
}

func NewUserService(db database.Users) *UserService {
	return &UserService{db: db}
}

func (u *UserService) AddNewUser(params *models.User) error {
	passHash, err := utils.Hash(params.Password)
	if err != nil {
		return err
	}
	params.Password = passHash
	return u.db.AddNewUser(params)
}

func (u *UserService) UpdateUserParams(params *models.User) error {
	return u.db.UpdateUserParams(params)
}

func (u *UserService) GetAllUsers(page, limit int) ([]*models.User, error) {
	offset := (page * limit) - limit
	return u.db.GetAllUsers(limit, offset)
}

func (u *UserService) GetUserById(userId int) (*models.User, error) {
	return u.db.GetUserById(userId)
}
func (u *UserService) GetUserByPhone(phone string) (*models.User, error) {
	user, err := u.db.GetUserByPhone(phone)
	if err != nil {
		return nil, err
	}
	passDeHash, _ := utils.DeHash(user.Password)
	user.Password = passDeHash
	return user, nil
}
func (u *UserService) DeleteUser(userId int) error {
	return u.db.DeleteUser(userId)
}
