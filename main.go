package main

import (
	"github.com/gin-api/modules/informations/informationsHandler"
	"github.com/gin-api/modules/informations/informationsRepository"
	"github.com/gin-api/modules/informations/informationsService"
	"github.com/gin-api/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	informationsDB := database.InitDatabaseMariaDB()

	infoRepository := informationsRepository.NewInformationsRepositoryDB(informationsDB)
	infoService := informationsService.NewInformationsService(&infoRepository)
	informationsHandler := informationsHandler.NewInformationsHandler(&infoService)

	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/informations", informationsHandler.GetInformations)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
