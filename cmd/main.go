package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/omegabytes/user_list/service"
)

var config struct {
	ServiceAddr string
}

func init() {
	config.ServiceAddr = ":3000"
}

func main() {
	// init the service
	svc := service.NewService(service.ServiceConfig{})

	// init handler
	handler := svc.NewServiceHandler()

	// init server
	server := &http.Server{
		Addr:         config.ServiceAddr,
		Handler:      handler,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// listen and serve
	go func() {
		fmt.Println("listening on: ", config.ServiceAddr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	shutdown(server)
}

func shutdown(server *http.Server) {
	osSignal := make(chan os.Signal, 1)
	signal.Notify(osSignal, syscall.SIGTERM, syscall.SIGINT)

	sig := <-osSignal

	ctx, cancel := context.WithTimeout(context.Background(), server.IdleTimeout)
	defer cancel()

	fmt.Println("Shutting down with signal: ", sig)
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("stopping http server with error ", err)
	}
}
