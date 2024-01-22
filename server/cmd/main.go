package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	cardmarketplace "github.com/AlekseySmirnovEmpire/CardMarketplace"
	"github.com/AlekseySmirnovEmpire/CardMarketplace/package/handler"
	"github.com/AlekseySmirnovEmpire/CardMarketplace/package/repository"
	"github.com/AlekseySmirnovEmpire/CardMarketplace/package/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("Ошибка чтения .env файла: '%s'", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{ConnectionString: os.Getenv("POSTGRES_CONNECTION_STRING")})
	if err != nil {
		logrus.Fatalf("Невозможно подключиться к БД: '%s'", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(cardmarketplace.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Ошибка работы приложение, сервер будет закрыт: '%s'", err.Error())
		}
	}()

	logrus.Print("Server started!")

	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGTERM, syscall.SIGINT)

	<-quite
	logrus.Print("Server shutting down ...")

	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("Error occured on server shutting down: '%s'", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured on server while closing DB connection: '%s'", err.Error())
	}

	logrus.Print("Server shut down!")
}

func initConfig() error {
	return godotenv.Load()
}
