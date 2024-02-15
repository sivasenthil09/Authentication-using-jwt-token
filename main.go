package main

import (
	"authentication/config"
	"authentication/constants"
	"authentication/controllers"
	"authentication/routes"
	services "authentication/service"

	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}
func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	collection := mongoClient.Database(constants.DatabaseName).Collection("senthil")
	ser := services.NewUserServ(collection, ctx)
	cont := controllers.New(ser)
	routes.RegisterUserRoutes(server, cont)
}

func main() {

	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
