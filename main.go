package main

import (
	"github.com/dwibedis/chat-service/app"
	"github.com/dwibedis/chat-service/app/constants"
	"github.com/dwibedis/chat-service/app/controllers"
	"github.com/dwibedis/chat-service/app/repository"
	"github.com/dwibedis/chat-service/app/service"
	"github.com/dwibedis/chat-service/app/validator"
	"github.com/dwibedis/chat-service/mongodb"
	"github.com/gorilla/mux"
	"log"
)

func main()  {
	mongoDbClient, err := mongodb.InitMongoDb()
	if err != nil {
		log.Println("Server bootstrap failed!!!!!")
		return
	}

	router := app.NewRouter(mux.NewRouter().StrictSlash(true))

	userRepo := repository.NewUserRepo(mongoDbClient)
	userOtpRepo := repository.NewUserOtp(mongoDbClient)
	userValidator := validator.NewUser(userRepo)
	userService := service.NewUser(userRepo, userValidator, userOtpRepo)
	userRegController := controllers.NewUserRegister(userService)
	router.RegisterNewHandler("/register", userRegController.RegisterUser, constants.METHOD_POST)
	err = router.ListenAndServe(8080)
	if err != nil {
		log.Println("Failed!!!!!!")
	}
	log.Println("Successful!!!!!")
}