package models

import (
	"gorm.io/gorm"
)

type Bet struct {
	gorm.Model
	Team1  string
	Team2  string
	Odd1   float64
	Odd2   float64
	Result int32
	IsOpen bool
}
