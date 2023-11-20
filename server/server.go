package server

import (
	"context"

	"github.com/superjcd/calendarservice/conf"
	v1 "github.com/superjcd/calendarservice/genproto/v1"
	"github.com/superjcd/calendarservice/service"
	"github.com/superjcd/calendarservice/service/implement"
	"github.com/superjcd/webdev_tookit/database"
	"gorm.io/gorm"
)

var _DB *gorm.DB

type GrpcServer struct {
	v1.UnimplementedCalendarServiceServer
	service service.CalendarService
	client  v1.CalendarServiceClient
	conf    *conf.Config
}

func NewServer(conf *conf.Config, client v1.CalendarServiceClient) (v1.CalendarServiceServer, error) {
	db, err := database.NewMysqlDb(conf.Db.Username, conf.Db.Password, conf.Db.Host, "3306", conf.Db.Database, &gorm.Config{})

	if err != nil {
		panic("Initialize database failed")
	}

	service, err2 := implement.NewCalendarService(db)

	if err2 != nil {
		panic("Create calendarservice failed")
	}

	server := &GrpcServer{
		service: service,
		client:  client,
		conf:    conf,
	}

	return server, nil
}

func (s *GrpcServer) CreateCalendarItem(ctx context.Context, rq *v1.CreateCalendarItemRequest) (*v1.CreateCalendarItemResponse, error) {
	err := s.service.CreateCalendarItem(ctx, rq)
	if err != nil {
		return &v1.CreateCalendarItemResponse{Msg: "failed", Status: v1.Status_failure}, err
	}

	return &v1.CreateCalendarItemResponse{Msg: "success", Status: v1.Status_success}, nil
}

func (s *GrpcServer) ListCalendarItems(ctx context.Context, rq *v1.ListCalendarItemsRequest) (*v1.ListCalendarItemsResponse, error) {
	resp, err := s.service.ListCalendarItems(ctx, rq)
	if err != nil {
		return &v1.ListCalendarItemsResponse{Msg: "failed", Status: v1.Status_failure}, err
	}

	response_items := make([]*v1.CalendarItem, 10)

	for _, item := range resp.Items {
		response_items = append(response_items, &v1.CalendarItem{
			Creator: item.Creator,
			Date:    item.Date,
			Content: item.Content,
		})
	}

	return &v1.ListCalendarItemsResponse{Msg: "success", Status: v1.Status_success, Items: response_items}, nil
}

func (s *GrpcServer) UpdateCalendarItem(ctx context.Context, rq *v1.UpdateCalendarItemRequest) (*v1.UpdateCalendarItemResponse, error) {
	err := s.service.UpdateCalendarItem(ctx, rq)

	if err != nil {
		return &v1.UpdateCalendarItemResponse{Msg: "failed", Status: v1.Status_failure}, err
	}
	return &v1.UpdateCalendarItemResponse{Msg: "success", Status: v1.Status_success}, nil
}
