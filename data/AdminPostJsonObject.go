package data

type AdminViewReportResponseJsonObject struct {
	ReportList any `json:"report_list"`
}

type AdminViewReportItem struct {
	PostId   int64  `json:"post_id"`
	Reason   string `json:"reason"`
	Content  string `json:"content"`
	UserName string `json:"username"`
	ReportId int64  `json:"report_id"`
}

type AuditedReportRequestJsonObject struct {
	UserId   int64 `json:"user_id"`
	ReportId int64 `json:"report_id"`
	Approval int   `json:"approval"`
}
