package models

type Pin struct {
	Id     int64  `json:"id"`
	Pin    string `json:"pin"`
	Status string `json:"status" gorm:"Default:'pending'"`
}
