package main

import (
	"fmt"
	"go-phone-book/config"
	"go-phone-book/infrastructure/dbconnection"
	"go-phone-book/src/application/createService"
	"go-phone-book/src/interface/createController"
	"go-phone-book/src/repository"

	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading env file: %s", err)
	}

	var config config.Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error on unmarshalling config: %s", err)
	}

	dbConnection, err := dbconnection.CreateDBConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	repository := repository.NewRepository(dbConnection)
	service := createService.NewService(*repository)
	controller := createController.NewController(*service)

	fmt.Println(controller)
	// http.HandleFunc("/", repository.Create(dbConnection))
}
