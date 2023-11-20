package service

import (
	"gorm.io/gorm"
)

//	First  string `gorm:"uniqueIndex:idx_first_second"`
//
// Second string `gorm:"uniqueIndex:idx_first_second"`
type CalendarItem struct {
	gorm.Model
	Date    string `json:"date" gorm:"column:date;uniqueIndex:idx_creator_date;size:20"`
	Creator string `json:"creator" gorm:"column:creator;uniqueIndex:idx_creator_date;size:30"`
	Content string `json:"content" gorm:"column:content"`
}

type ListCalendarItemResponse struct {
	TotalCount int64           `json:"totalCount"`
	Items      []*CalendarItem `json:"items"`
}
