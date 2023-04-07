package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"restapi/config"
	"restapi/internal/server"
	"restapi/internal/usecase"

	"restapi/internal/controller"
)

func Start(cnf *config.Config) error {
	usecase := usecase.NewUsecase()
	handler := controller.NewHandler(usecase)
	router := controller.RoutesInit(handler)
	server := server.New(cnf, router)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	server.Start()

	select {
	case s := <-interrupt:
		log.Printf("signal: " + s.String())
	case err := <-server.Notify():
		log.Printf("signal.Notify: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		log.Printf("server shutdown: %v", err)
		return err
	}
	log.Printf("app: graceful shutdown")
	return nil
}
