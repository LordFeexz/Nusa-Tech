package models

type Balance struct {
	Id         int64 `json:"id" gorm:"primaryKey"`
	CurrencyId int64 `json:"CurrencyId" gorm:"not null;ForeignKey:Id"`
	UserId     int64 `json:"UserId" gorm:"not null;ForeignKey:Id"`
	Amount     int64 `json:"amount" gorm:"DEFAULT:0"`
}
