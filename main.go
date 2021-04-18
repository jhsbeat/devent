package main

import (
	"devent/controllers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	r.LoadHTMLGlob("templates/**/*")

	eventRepo := controllers.New()
	r.GET("/", eventRepo.GetEvents)

	return r
}
