//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/superjcd/calendarservice/conf"
	v1 "github.com/superjcd/calendarservice/genproto/v1"
	"github.com/superjcd/calendarservice/server"
)

// InitServer Inject service's component
func InitServer(conf *conf.Config) (v1.CalendarServiceServer, error) {

	wire.Build(
		server.NewClient,
		server.NewServer,
	)

	return &server.GrpcServer{}, nil

}
