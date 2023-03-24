package transactionsdto

import (
	"dumbflix/models"
	"time"
)

type TransactionRequest struct {
	ID        int                         `json:"id" gorm:"primary_key:auto_increment"`
	StartDate time.Time                   `json:"startdate" form:"startdate" `
	DueDate   time.Time                   `json:"duedate" form:"duedate" `
	User      models.UsersProfileResponse `json:"userId"`
	UserID    int                         `json:"user_id"`
	Attache   string                      `json:"attach" form:"image" gorm:"type: text" `
	Status    string                      `json:"status" form:"status" gorm:"type: text"`
}
