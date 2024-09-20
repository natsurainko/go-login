package dao

import (
	"login-project/data"

	"gorm.io/gorm"
)

type ReportDao struct {
	DataBase *gorm.DB
}

func (dao ReportDao) AddReport(reason string, userId int64, postId int64) error {
	result := dao.DataBase.Create(&data.ReportDataRecord{
		Reason: reason,
		UserId: userId,
		PostId: postId,
		Status: 0,
	})

	return result.Error
}

func (dao ReportDao) GetReportsFromUserId(userId int64) ([]data.ReportDataRecord, error) {
	var reportsOfUser []data.ReportDataRecord
	result := dao.DataBase.Find(&reportsOfUser, "user_id = ?", userId)

	return reportsOfUser, result.Error
}

func (dao ReportDao) GetUnauditedReports() ([]data.ReportDataRecord, error) {
	var reportsDatas []data.ReportDataRecord
	result := dao.DataBase.Find(&reportsDatas, "status = ?", 0)

	return reportsDatas, result.Error
}

func (dao ReportDao) FindReportByReportId(reportId int64) (data.ReportDataRecord, error) {
	var reportData data.ReportDataRecord
	result := dao.DataBase.First(&reportData, "report_id = ?", reportId)

	return reportData, result.Error
}

func (dao ReportDao) UpdateReportStatus(reportId int64, status int) error {
	result := dao.DataBase.Model(data.ReportDataRecord{}).Where("report_id = ?", reportId).Update("Status", status)

	return result.Error
}
