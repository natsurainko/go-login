package services

import (
	"login-project/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DataBaseService struct {
	DataBase *gorm.DB
}

func (service *DataBaseService) InitDataBase() error {
	db, error := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if error != nil {
		return error
	}

	service.DataBase = db

	if !db.Migrator().HasTable(&data.UserDataRecord{}) {
		db.AutoMigrate(&data.UserDataRecord{})
	}
	if !db.Migrator().HasTable(&data.PostDataRecord{}) {
		db.AutoMigrate(&data.PostDataRecord{})
	}
	if !db.Migrator().HasTable(&data.ReportDataRecord{}) {
		db.AutoMigrate(&data.ReportDataRecord{})
	}

	return nil
}
