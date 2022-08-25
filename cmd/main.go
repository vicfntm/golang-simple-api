package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"github.com/vicfntm/splitshit2"
	"github.com/vicfntm/splitshit2/src/conn"
	"github.com/vicfntm/splitshit2/src/handlers"
	"github.com/vicfntm/splitshit2/src/repositories"
	"github.com/vicfntm/splitshit2/src/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("env file loading failed: %s", err.Error())
	}

	dbConn, err := conn.NewPostgresConn(conn.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	})

	if err != nil {
		log.Fatalf("db connection has been lost: %s", err.Error())
	}
	repositories := repositories.NewRepository(dbConn)
	services := services.NewService(repositories)
	srv := new(splitshit2.Server)
	handler := handlers.NewHandler(services)

	if error := initConfig(); error != nil {
		log.Fatalf("configs has not work properly: %s", error.Error())
	}
	if err := srv.Run(os.Getenv("SERVER_PORT"), handler.InitRoutes()); err != nil {
		log.Fatalf("error while server running %s", err.Error())
	}
	fmt.Printf("run on port: %s", os.Getenv("SERVER_PORT"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
