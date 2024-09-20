package services

import (
	"login-project/dao"
	"login-project/data"
)

type PostService struct {
	PostDao dao.PostDao
}

func (service PostService) ReleasePost(requestBody data.ReleasePostRequestJsonObject) error {
	return service.PostDao.AddPost(requestBody.Content, requestBody.UserId)
}

func (service PostService) FetchAllPosts() ([]data.PostDataRecord, error) {
	return service.PostDao.GetAllPosts()
}

func (service PostService) FindPostByPostId(postId int64) (data.PostDataRecord, error) {
	return service.PostDao.FindPostByPostId(postId)
}

func (service PostService) UpdatePostContent(content string, postId int64) error {
	return service.PostDao.UpdatePostContent(content, postId)
}

func (service PostService) DeletePost(postId int64) error {
	return service.PostDao.DeletePost(postId)
}
