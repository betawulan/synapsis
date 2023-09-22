package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/betawulan/synapsis/delivery"
	"github.com/betawulan/synapsis/repository"
	"github.com/betawulan/synapsis/service"
)

func main() {
	viper.AutomaticEnv()
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("failed running because file .env")
	}

	dsn := viper.GetString("mysql_dsn")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("can't connect database")
	}

	secretKey := viper.GetString("secret_key")
	port := viper.GetString("port")

	authRepo := repository.NewAuthRepository(db)
	onlineStoreRepo := repository.NewOnlineStoreRepository(db)

	authService := service.NewAuthService(authRepo, []byte(secretKey))
	onlineStoreService := service.NewOnlineStoreService(onlineStoreRepo)

	e := echo.New()
	delivery.AddAuthRoute(authService, e)
	delivery.RegisterOnlineStoreRoute(onlineStoreService, e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
