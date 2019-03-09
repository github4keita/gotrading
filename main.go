package main

import (
	"fmt"
	"gotrading/app/models"
	"time"
)

// import (
// 	"gotrading/app/controllers"
// 	"gotrading/config"
// 	"gotrading/utils"
// )

// func main() {
// 	utils.LoggingSetting(config.Config.LogFile)
// 	controllers.StreamIngestionData()
// 	controllers.StartWebServer()
// }

func main() {
	s := models.NewSignalEvents()
	df, _ := models.GetAllCandle("BTC_JPY", time.Minute, 10)
	c1 := df.Candles[0]
	c2 := df.Candles[5]
	s.Buy("BTC_JPY", c1.Time.UTC(), c1.Close, 1.0, true)
	s.Sell("BTC_JPY", c2.Time.UTC(), c2.Close, 1.0, true)
	fmt.Println(s.CollectAfter(time.Now().UTC()))
	fmt.Println(s.CollectAfter(c1.Time))
}
