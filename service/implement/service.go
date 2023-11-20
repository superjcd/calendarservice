package implement

import (
	"context"
	"fmt"
	"sync"

	v1 "github.com/superjcd/calendarservice/genproto/v1"
	"github.com/superjcd/calendarservice/service"
	"gorm.io/gorm"
)

type clandar_service struct {
	db *gorm.DB
}

var (
	cs   service.CalendarService
	once sync.Once
)

func NewCalendarService(db *gorm.DB) (service.CalendarService, error) {
	if db == nil && cs == nil {
		return nil, fmt.Errorf("failed to get a bew calndar service instance")
	}

	once.Do(func() {
		db.AutoMigrate(&service.CalendarItem{})
		cs = &clandar_service{db: db}
	})

	return cs, nil
}

func (s *clandar_service) CreateCalendarItem(ctx context.Context, rq *v1.CreateCalendarItemRequest) error {
	item := service.CalendarItem{
		Creator: rq.Creator,
		Date:    rq.Date,
		Content: rq.Content,
	}

	return s.db.Create(&item).Error
}

func (s *clandar_service) ListCalendarItems(ctx context.Context, rq *v1.ListCalendarItemsRequest) (service.ListCalendarItemResponse, error) {
	var resp service.ListCalendarItemResponse

	tx := s.db

	if rq.Creator != "" {
		tx = tx.Where("creator = ?", rq.Creator)
	}

	d := tx.
		Find(&resp.Items).
		Offset(-1).
		Limit(-1).
		Count(&resp.TotalCount)

	return resp, d.Error
}

func (s *clandar_service) UpdateCalendarItem(ctx context.Context, rq *v1.UpdateCalendarItemRequest) error {
	item := service.CalendarItem{}

	if err := s.db.Where("date = ? and creator = ?", rq.Date, rq.Creator).First(&item).Error; err != nil {
		return err
	}

	if rq.Content != "" {
		item.Content = rq.Content
	}

	return s.db.Save(&item).Error
}

func (s *clandar_service) Close() error {
	db, _ := s.db.DB()

	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
