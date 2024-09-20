package data

type RegisterRequestJsonObject struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	UserType int    `json:"user_type"`
}
