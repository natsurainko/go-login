package data

type UserDataRecord struct {
	UserName string
	Name     string
	Password string
	UserType int
	UserId   int64 `gorm:"uniqueIndex;autoIncrement"`
}

type PostDataRecord struct {
	Content   string
	UserId    int64
	PostId    int64 `gorm:"uniqueIndex;autoIncrement"`
	Time      string
	IsDeleted bool `default:"false" json:"-"`
}

type ReportDataRecord struct {
	UserId   int64
	PostId   int64
	Reason   string
	Status   int
	ReportId int64 `gorm:"uniqueIndex;autoIncrement"`
}
