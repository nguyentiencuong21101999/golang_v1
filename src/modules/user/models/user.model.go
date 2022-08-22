package UserModels

type User struct {
	UserId   uint   `json:"userId" gorm:"primary_key;autoIncrement;column:UserId"`
	UserName string `json:"userName" gorm:"column:UserName"`
	Password string `json:"password" gorm:"column:Password"`
}
