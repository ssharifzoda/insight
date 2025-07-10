package service

import (
	"insight/internal/database"
	"insight/internal/models"
)

type UserService struct {
	db database.Users
}

func NewUserService(db database.Users) *UserService {
	return &UserService{db: db}
}

func (u *UserService) AddNewUser(params *models.User) error {
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
	return u.db.GetUserByPhone(phone)
}
func (u *UserService) DeleteUser(userId int) error {
	return u.db.DeleteUser(userId)
}
