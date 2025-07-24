package service

import (
	"github.com/google/uuid"
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

func (u *UserService) AddNewUser(params *models.User) (*models.User, error) {
	pass, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	passHash, err := utils.Hash(pass.String())
	if err != nil {
		return nil, err
	}
	params.Password = passHash
	userResponse, err := u.db.AddNewUser(params)
	if err != nil {
		return nil, err
	}
	userResponse.Password, err = utils.DeHash(userResponse.Password)
	if err != nil {
		return nil, err
	}
	return userResponse, nil
}

func (u *UserService) UpdateUserParams(params *models.User) error {
	return u.db.UpdateUserParams(params)
}

func (u *UserService) GetAllUsers(page, limit int) ([]*models.User, int, error) {
	offset := (page * limit) - limit
	return u.db.GetAllUsers(limit, offset)
}

func (u *UserService) GetUserById(userId int) (*models.UserInfo, error) {
	item, err := u.db.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return item, nil
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
