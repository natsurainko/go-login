package dao

import (
	"login-project/data"
	"time"

	"gorm.io/gorm"
)

type PostDao struct {
	DataBase *gorm.DB
}

func (dao PostDao) AddPost(content string, userId int64) error {
	result := dao.DataBase.Create(&data.PostDataRecord{
		Content: content,
		UserId:  userId,
		Time:    time.Now().String(),
	})

	return result.Error
}

func (dao PostDao) GetAllPosts() ([]data.PostDataRecord, error) {
	var postDataRecords []data.PostDataRecord
	result := dao.DataBase.Find(&postDataRecords, "is_deleted = ?", false)

	return postDataRecords, result.Error
}

func (dao PostDao) FindPostByPostId(postId int64) (data.PostDataRecord, error) {
	var postData data.PostDataRecord
	result := dao.DataBase.First(&postData, "post_id = ?", postId)

	return postData, result.Error
}

func (dao PostDao) UpdatePostContent(content string, postId int64) error {
	result := dao.DataBase.Model(data.PostDataRecord{}).Where("post_id = ?", postId).Update("Content", content)
	return result.Error
}

func (dao PostDao) DeletePost(postId int64) error {
	result := dao.DataBase.Model(data.PostDataRecord{}).Where("post_id = ?", postId).Update("IsDeleted", true)
	return result.Error
}
