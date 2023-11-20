package service

import (
	"context"

	v1 "github.com/superjcd/calendarservice/genproto/v1"
)

type CalendarService interface {
	CreateCalendarItem(ctx context.Context, rq *v1.CreateCalendarItemRequest) error
	ListCalendarItems(ctx context.Context, rq *v1.ListCalendarItemsRequest) (ListCalendarItemResponse, error)
	UpdateCalendarItem(ctx context.Context, rq *v1.UpdateCalendarItemRequest) error
	Close() error
}
