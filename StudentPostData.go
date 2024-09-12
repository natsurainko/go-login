package main

type ReleasePostRequestJsonObject struct {
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
}

type FetchAllPostsResponseJsonObject struct {
	PostList []PostDataRecord `json:"post_list"`
}

type ModifyPostRequestJsonObject struct {
	UserId  int64  `json:"user_id"`
	Content string `json:"content"`
	PostId  int64  `json:"post_id"`
}

type ReportPostRequestJsonObject struct {
	UserId int64  `json:"user_id"`
	PostId int64  `json:"post_id"`
	Reason string `json:"reason"`
}

type ViewReportResponseJsonObject struct {
	ReportList any `json:"report_list"`
}

type ViewReportItem struct {
	PostId  int64  `json:"post_id"`
	Reason  string `json:"reason"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}
