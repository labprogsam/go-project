package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Password  string
	Username  string
	IsAdmin   bool
	Pontuacao float64
}
