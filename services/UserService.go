package services

import (
	"errors"
	"login-project/dao"
	"login-project/data"

	"gorm.io/gorm"
)

type UserService struct {
	UserDao dao.UserDao
}

func (service UserService) RegisterUser(requestBody data.RegisterRequestJsonObject) error {
	return service.UserDao.AddUser(requestBody.Name, requestBody.UserName, requestBody.Password, requestBody.UserType)
}

func (service UserService) FindUserByUserName(userName string) (data.UserDataRecord, error) {
	return service.UserDao.FindUserByUserName(userName)
}

func (service UserService) FindUserByUserId(userId int64) (data.UserDataRecord, error) {
	return service.UserDao.FindUserByUserId(userId)
}

func (service UserService) ContainsUserName(userName string) (bool, error) {
	user, error := service.UserDao.FindUserByUserName(userName)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user.UserName == userName, error
}

func (service UserService) ContainsUserId(userId int64) (bool, error) {
	user, error := service.UserDao.FindUserByUserId(userId)
	if errors.Is(error, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user.UserId == userId, error
}
