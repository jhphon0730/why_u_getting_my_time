package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jhphon0730/action_manager/internal/config"
	"github.com/jhphon0730/action_manager/internal/database"
	"github.com/jhphon0730/action_manager/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// config 설정 초기화
	config, err := config.LoadConfig()
	if err != nil || config == nil {
		log.Fatalln("Failed to load Config", err.Error())
	}

	// 데이터베이스 연결
	db, err := database.NewDB()
	if err != nil || db == nil {
		log.Fatalln("Failed to connect to Database", err.Error())
	}
	if err := database.AutoMigrate(db.DB); err != nil {
		log.Fatalln("Failed to migrate Database", err.Error())
	}
	sqlDB, _ := db.DB.DB()
	defer sqlDB.Close()

	// 서버 옵션 설정
	PORT := config.PORT
	APP_ENV := config.APP_ENV

	srv := server.NewServer(PORT, APP_ENV)
	srv.RegisterRoutes()

	go func() {
		log.Println("Server Running PORT : ", PORT)
		if err := srv.Start(); err != nil {
			log.Fatalln("Failed to run server", err.Error())
		}
	}()

	<-ctx.Done()

	log.Println("Shutting Down...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.ShutDown(shutdownCtx)
	log.Println("Server Stopped ")
}
