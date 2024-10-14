package models

import (
	"gorm.io/gorm"
)

type UserBet struct {
	gorm.Model
	UserID uint
	BetID  uint
	Team   int
}
