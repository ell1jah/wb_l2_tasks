package main

import (
	"context"
	"d11/config"
	"d11/database"
	"d11/handler"
	server2 "d11/server"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	db := database.NewDataBase()
	handler := handler.NewHandler(db)
	router := handler.InitRoutes()

	server := new(server2.Server)
	go func() {
		if err := server.Run(cfg, router); err != nil {
			log.Fatalf("error running http server: %s", err.Error())
		}
	}()
	log.Print("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("app stopped")

	if err := server.ShutDown(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "error occured on server shutting down: %s", err.Error())
	}
}
