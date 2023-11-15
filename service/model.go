package service

import (
	"gorm.io/gorm"
)

type CalendarItem struct {
	gorm.Model
	Date    string `json:"date" group:"date"`
	Creator string `json:"creator" gorm:"column:creator"`
	Content string `json:"content" group:"content"`
}
