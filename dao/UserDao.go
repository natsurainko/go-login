package dao

import (
	"login-project/data"

	"gorm.io/gorm"
)

type UserDao struct {
	DataBase *gorm.DB
}

func (dao UserDao) AddUser(name string, userName string, password string, userType int) error {
	result := dao.DataBase.Create(&data.UserDataRecord{
		Name:     name,
		UserName: userName,
		Password: password,
		UserType: userType,
	})

	return result.Error
}

func (dao UserDao) FindUserByUserName(userName string) (data.UserDataRecord, error) {
	var userData data.UserDataRecord
	result := dao.DataBase.First(&userData, "user_name = ?", userName)

	return userData, result.Error
}

func (dao UserDao) FindUserByUserId(userId int64) (data.UserDataRecord, error) {
	var userData data.UserDataRecord
	result := dao.DataBase.First(&userData, "user_id = ?", userId)

	return userData, result.Error
}
