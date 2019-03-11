package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
	"log"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	// controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
