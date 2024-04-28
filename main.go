package main

import (
	"go-phone-book/config"
	"go-phone-book/infrastructure/dbconnection"
	"go-phone-book/src/application/user/create/createService"
	"go-phone-book/src/interface/user/create"
	"go-phone-book/src/repository"
	"net/http"

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
	service := createService.NewService(repository)
	controller := createUser.NewController(service)

	createUserController := createUser.UserCreateController(*controller)

	http.HandleFunc("/create-user", createUserController.CreateUser)

	http.ListenAndServe(":8080", nil)
}
