package server

import (
	"context"
	"fmt"
	"time"

	"github.com/superjcd/calendarservice/conf"
	v1 "github.com/superjcd/calendarservice/genproto/v1"
)

// Run Run service server
func Run(cfg string) {

	conf := conf.NewConfig(cfg)
	// run grpc server
	RunGrpcServer(initServer(conf), conf)
	// listen exit server event
	HandleExitServer(conf)

}

// SetServer Wire inject service's component
func initServer(conf *conf.Config) v1.CalendarServiceServer {
	server, err := InitServer(conf)
	if err != nil {
		panic("run server failed.[ERROR]=>" + err.Error())
	}
	return server
}

// HandleExitServer Handle service exit event
func HandleExitServer(conf *conf.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conf.Grpc.Server.GracefulStop()
	<-ctx.Done()
	fmt.Println("Graceful shutdown http & grpc server.")

}
