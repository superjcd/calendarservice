package server

import (
	"context"
	"time"

	"github.com/superjcd/calendarservice/conf"
	v1 "github.com/superjcd/calendarservice/genproto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient(conf *conf.Config) (v1.CalendarServiceClient, error) {

	serverAddress := conf.Grpc.Host + conf.Grpc.Port
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := v1.NewCalendarServiceClient(conn)
	return client, nil

}
