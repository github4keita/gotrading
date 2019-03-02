package main

import (
	"gotrading/app/controllers"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSetting(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
