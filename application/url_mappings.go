package application

import (
	"github.com/hariprathap-hp/tinyURL_KGS_API/controller/keys"
	"github.com/hariprathap-hp/tinyURL_KGS_API/controller/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/populate", keys.PopulateKeys)
	router.GET("/getKeys", keys.GetKeys)
}
