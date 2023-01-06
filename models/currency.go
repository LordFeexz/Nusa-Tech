package models

type Currency struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Currency string `json:"currency"`
}
