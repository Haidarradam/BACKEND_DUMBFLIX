package models

import "time"

type Transaction struct {
	ID        int                  `json:"id" gorm:"primary_key:auto_increment"`
	Startdate time.Time            `json:"startdate"`
	Duedate   time.Time            `json:"duedate"`
	Attache   string               `json:"attach" gorm:"type: text"`
	Status    string               `json:"status" gorm:"type: text"`
	User      UsersProfileResponse `json:"user"`
	UserID    int                  `json:"user_id"`
	Price     int                  `json:"price"`
}

type TransactionInUser struct {
	ID        int                  `json:"id"`
	Startdate time.Time            `json:"startdate"`
	Duedate   time.Time            `json:"duedate"`
	Attache   string               `json:"attach" gorm:"type: text"`
	Status    string               `json:"status" gorm:"type: text"`
	User      UsersProfileResponse `json:"user"`
	UserID    int                  `json:"user_id"`
}

func (TransactionInUser) TableName() string {
	return "transactions"
}
