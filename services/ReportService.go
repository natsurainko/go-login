package services

import (
	"login-project/dao"
	"login-project/data"

	"github.com/ahmetb/go-linq/v3"
)

type ReportService struct {
	ReportDao   dao.ReportDao
	UserService UserService
	PostService PostService
}

func (service ReportService) AddReport(requestBody data.ReportPostRequestJsonObject) error {
	return service.ReportDao.AddReport(requestBody.Reason, requestBody.UserId, requestBody.PostId)
}

func (service ReportService) FindReportByReportId(reportId int64) (data.ReportDataRecord, error) {
	return service.ReportDao.FindReportByReportId(reportId)
}

func (service ReportService) ViewReportsFromUserId(userId int64) (data.ViewReportResponseJsonObject, error) {
	reportsOfUser, error := service.ReportDao.GetReportsFromUserId(userId)

	var ReportList []interface{} = linq.From(reportsOfUser).SelectT(func(report data.ReportDataRecord) data.ViewReportItem {
		post, _ := service.PostService.FindPostByPostId(report.PostId)

		return data.ViewReportItem{
			PostId:  report.PostId,
			Reason:  report.Reason,
			Content: post.Content,
			Status:  report.Status,
		}
	}).Results()

	var DefaultReportList [0]data.ViewReportItem

	if len(ReportList) == 0 {
		return data.ViewReportResponseJsonObject{
			ReportList: DefaultReportList,
		}, error
	} else {
		return data.ViewReportResponseJsonObject{
			ReportList: ReportList,
		}, error
	}
}

func (service ReportService) AdminViewReports() (data.AdminViewReportResponseJsonObject, error) {
	reportsDatas, error := service.ReportDao.GetUnauditedReports()

	var ReportList []interface{} = linq.From(reportsDatas).SelectT(func(report data.ReportDataRecord) data.AdminViewReportItem {
		user, _ := service.UserService.FindUserByUserId(report.UserId)
		post, _ := service.PostService.FindPostByPostId(report.PostId)

		return data.AdminViewReportItem{
			PostId:   report.PostId,
			Reason:   report.Reason,
			Content:  post.Content,
			UserName: user.UserName,
			ReportId: report.ReportId,
		}
	}).Results()

	var DefaultReportList [0]data.ViewReportItem

	if len(ReportList) == 0 {
		return data.AdminViewReportResponseJsonObject{
			ReportList: DefaultReportList,
		}, error
	} else {
		return data.AdminViewReportResponseJsonObject{
			ReportList: ReportList,
		}, error
	}
}

func (service ReportService) AuditedReport(reportId int64, postId int64, approval int) error {
	error := service.ReportDao.UpdateReportStatus(reportId, approval)

	if error == nil && approval == 1 {
		return service.PostService.DeletePost(postId)
	}

	return error
}
