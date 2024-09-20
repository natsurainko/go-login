package data

type LoginRequestJsonObject struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginDataResponseJsonObject struct {
	UserId   int64 `json:"user_id"`
	UserType int   `json:"user_type"`
}
