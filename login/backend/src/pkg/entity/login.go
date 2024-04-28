package entity

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	User  *User
	Token string
}
