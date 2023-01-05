package models

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"varchar(30)"`
	Password string `json:"password"`
	Status   string `json:"status" gorm:"DEFAULT:'pending'"`
}
