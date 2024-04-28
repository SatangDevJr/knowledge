package entity

type User struct {
	ID              int64  `json:"id" sql:"id"`
	UserName        string `json:"username" sql:"username"`
	Password        string `json:"password" sql:"password"`
	CreatedByUserId int64  `json:"createdbyuserid" sql:"createdbyuserid"`
	UpdatedByUserId int64  `json:"updatedbyuserid" sql:"updatedbyuserid"`
	DelFlag         bool   `json:"delflag" sql:"delflag"`
}
